package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLength  = 6 // Fixed length for the short code
)

// Add these new types to track user information and URL history
type UserInfo struct {
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Browser   string    `json:"browser"`
	OS        string    `json:"os"`
	Device    string    `json:"device"`
	CreatedAt time.Time `json:"created_at"`
}

type URLCreation struct {
	ShortCode       string
	LongURL         string
	CreatedAt       time.Time
	UserInfo        UserInfo
	ViewCount       int // Add this field
	UniqueViewCount int // Add this field
}

// Modify URLShortener struct to include user tracking
type URLShortener struct {
	mu          sync.RWMutex
	store       map[string]*URLData
	userHistory map[string][]URLCreation // IP -> URLs created by user
	domain      string
}

// URLData holds the original long URL and view metrics.
type URLData struct {
	LongURL     string
	ViewCount   uint64
	UniqueViews map[string]bool
}

// shortenRequest and shortenResponse define the JSON request/response for shortening URLs.
type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// URLStats holds data to be displayed on the stats page.
type URLStats struct {
	LongURL         string
	ViewCount       uint64
	UniqueViewCount int
}

// Update NewURLShortener to initialize userHistory
func NewURLShortener(domain string) *URLShortener {
	return &URLShortener{
		store:       make(map[string]*URLData),
		userHistory: make(map[string][]URLCreation),
		domain:      domain,
	}
}

// Add function to parse user agent
func parseUserAgent(userAgent string) (browser, os, device string) {
	ua := strings.ToLower(userAgent)

	// Parse browser
	switch {
	case strings.Contains(ua, "firefox"):
		browser = "Firefox"
	case strings.Contains(ua, "chrome"):
		browser = "Chrome"
	case strings.Contains(ua, "safari"):
		browser = "Safari"
	case strings.Contains(ua, "edge"):
		browser = "Edge"
	default:
		browser = "Unknown"
	}

	// Parse OS
	switch {
	case strings.Contains(ua, "windows"):
		os = "Windows"
	case strings.Contains(ua, "mac"):
		os = "MacOS"
	case strings.Contains(ua, "linux"):
		os = "Linux"
	case strings.Contains(ua, "android"):
		os = "Android"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		os = "iOS"
	default:
		os = "Unknown"
	}

	// Parse device type
	switch {
	case strings.Contains(ua, "mobile"):
		device = "Mobile"
	case strings.Contains(ua, "tablet"):
		device = "Tablet"
	default:
		device = "Desktop"
	}

	return
}

// Update HandleShorten to track user information
func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Get user information
	ip := getIP(r)
	browser, os, device := parseUserAgent(r.UserAgent())
	userInfo := UserInfo{
		IP:        ip,
		UserAgent: r.UserAgent(),
		Browser:   browser,
		OS:        os,
		Device:    device,
		CreatedAt: time.Now(),
	}

	// Generate code and create URL data
	code := us.generateShortCode()
	data := &URLData{
		LongURL:     req.URL,
		ViewCount:   0,
		UniqueViews: make(map[string]bool),
	}

	// Store URL and update user history
	us.mu.Lock()
	us.store[code] = data
	urlCreation := URLCreation{
		ShortCode: code,
		LongURL:   req.URL,
		CreatedAt: time.Now(),
		UserInfo:  userInfo,
	}
	us.userHistory[ip] = append(us.userHistory[ip], urlCreation)
	us.mu.Unlock()

	shortURL := fmt.Sprintf("%s/%s", us.domain, code)
	resp := shortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// generateShortCode returns a base62-encoded string from a random integer.
func (us *URLShortener) generateShortCode() string {
	// Seed the random generator (ideally, do this once in your program's initialization)
	rand.Seed(time.Now().UnixNano())
	// Generate a random int64 value.
	num := rand.Int63()
	return encodeBase62(num)
}

// encodeBase62 converts a number to a base62 string and pads it to a fixed length.
func encodeBase62(num int64) string {
	if num == 0 {
		return fmt.Sprintf("%0*s", codeLength, "0")
	}

	var encoded []byte
	for num > 0 {
		remainder := num % 62
		encoded = append([]byte{base62Chars[remainder]}, encoded...)
		num /= 62
	}

	// Pad with zeros if necessary to ensure fixed length.
	if len(encoded) < codeLength {
		pad := make([]byte, codeLength-len(encoded))
		for i := range pad {
			pad[i] = '0'
		}
		encoded = append(pad, encoded...)
	}

	return string(encoded)
}

// HandleRedirect processes GET requests.
// If the path is "/" it serves the home page, otherwise it treats the path as a short code.
func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// If path is "/" serve the home page.
	if r.URL.Path == "/" {
		serveHomePage(w, r)
		return
	}

	// If the path starts with "/stats", let the stats handler take over.
	if strings.HasPrefix(r.URL.Path, "/stats") {
		us.HandleStats(w, r)
		return
	}

	// Otherwise, assume the path is a short code.
	code := r.URL.Path[1:]
	us.mu.Lock()
	data, exists := us.store[code]
	if exists {
		// Increment total view count.
		data.ViewCount++
		// Track unique views based on IP.
		ip := getIP(r)
		data.UniqueViews[ip] = true
	}
	us.mu.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, data.LongURL, http.StatusFound)
}

// HandleStats displays details for a given short code.
func (us *URLShortener) HandleStats(w http.ResponseWriter, r *http.Request) {
	// Expected URL: /stats/{code}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Stats code not provided", http.StatusBadRequest)
		return
	}
	code := parts[2]

	us.mu.RLock()
	data, exists := us.store[code]
	us.mu.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	stats := URLStats{
		LongURL:         data.LongURL,
		ViewCount:       data.ViewCount,
		UniqueViewCount: len(data.UniqueViews),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := statsTemplate.Execute(w, stats); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// getIP extracts the IP address from the request, considering the possibility of proxy headers.
func getIP(r *http.Request) string {
	// Check X-Forwarded-For header (if behind a proxy)
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	// Fallback to remote address.
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// serveHomePage renders the home page.
func serveHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	shortener := NewURLShortener(os.Getenv("DOMAIN"))

	// API endpoint to shorten URLs.
	http.HandleFunc("/shorten", shortener.HandleShorten)
	// /stats/{code} for URL statistics.
	http.HandleFunc("/stats/", shortener.HandleStats)
	// All other requests handled by HandleRedirect (home page or redirection).
	http.HandleFunc("/", shortener.HandleRedirect)
	// Add the new route in main()
	http.HandleFunc("/history", shortener.HandleHistory)
	http.HandleFunc("/delete/", shortener.HandleDelete)

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	fmt.Println("Server started at :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
