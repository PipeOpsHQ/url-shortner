package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

// URLData holds the original long URL and view metrics.
type URLData struct {
	LongURL     string
	ViewCount   uint64
	UniqueViews map[string]bool
}

// URLShortener holds the URL mappings and a counter for unique code generation.
type URLShortener struct {
	mu      sync.RWMutex
	store   map[string]*URLData
	counter uint64
	domain  string
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

// NewURLShortener initializes the URLShortener with a domain and an empty store.
func NewURLShortener(domain string) *URLShortener {
	return &URLShortener{
		store:  make(map[string]*URLData),
		domain: domain,
	}
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

// HandleShorten processes the POST /shorten endpoint to create a short URL.
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

	// Generate a unique short code and store the mapping.
	code := us.generateShortCode()
	data := &URLData{
		LongURL:     req.URL,
		ViewCount:   0,
		UniqueViews: make(map[string]bool),
	}

	us.mu.Lock()
	us.store[code] = data
	us.mu.Unlock()

	// Return the generated short URL.
	shortURL := fmt.Sprintf("%s/%s", us.domain, code)
	resp := shortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
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

// Home page template with animated gradients and iconic buttons
// Home page template with enhanced design
var homeTemplate = template.Must(template.New("home").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>URL Shortener</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/vue@2/dist/vue.js"></script>
    <style>
			@keyframes gradientBG {
			        0% { background-position: 0% 50%; }
			        50% { background-position: 100% 50%; }
			        100% { background-position: 0% 50%; }
			    }

        @keyframes float {
            0% { transform: translateY(0px); }
            50% { transform: translateY(-10px); }
            100% { transform: translateY(0px); }
        }

        @keyframes pulse {
            0%, 100% { transform: scale(1); }
            50% { transform: scale(1.05); }
        }

        @keyframes glow {
            0%, 100% { box-shadow: 0 0 5px rgba(59, 130, 246, 0.5); }
            50% { box-shadow: 0 0 20px rgba(59, 130, 246, 0.8); }
        }

	body {
        background: linear-gradient(-45deg, #0f172a, #1e3a8a, #0f172a, #1e3a8a);
        background-size: 400% 400%;
        animation: gradientBG 15s ease infinite;
        min-height: 100vh;
    }

	.card-gradient {
            background: linear-gradient(180deg, rgba(30, 41, 59, 0.9) 0%, rgba(15, 23, 42, 0.9) 100%);
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255, 255, 255, 0.1);
	}

	.icon-button {
        background: linear-gradient(45deg, #1e3a8a, #2563eb);
        transition: all 0.3s ease, background 0.5s ease;
        animation: glow 3s infinite;
    }

    .icon-button:hover {
        background: linear-gradient(45deg, #2563eb, #1e3a8a);
        background-size: 200% 200%;
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(59, 130, 246, 0.3);
    }

        .floating-card {
            animation: float 6s ease-in-out infinite;
        }

        .input-gradient {
            background: linear-gradient(90deg, rgba(30, 41, 59, 0.8) 0%, rgba(15, 23, 42, 0.8) 100%);
            backdrop-filter: blur(5px);
        }

        .github-link {
            transition: all 0.3s ease;
        }

        .github-link:hover {
            transform: translateY(-2px);
            filter: brightness(1.2);
        }

        .glow-text {
            text-shadow: 0 0 10px rgba(59, 130, 246, 0.5);
        }

        .particle {
            position: fixed;
            pointer-events: none;
            opacity: 0;
            background: white;
            border-radius: 50%;
        }

        @keyframes particle-animation {
            0% { transform: translate(0, 0); opacity: 1; }
            100% { transform: translate(var(--tx), var(--ty)); opacity: 0; }
        }
    </style>
</head>
<body class="antialiased">
    <div id="app" class="min-h-screen flex flex-col items-center justify-center p-4">
        <div class="max-w-2xl w-full floating-card">
            <h1 class="text-4xl font-bold mb-2 text-center text-blue-200 glow-text">URL Shortener</h1>
            <p class="text-gray-400 text-center mb-8">Create optimized short URLs for your links</p>

            <div class="card-gradient rounded-xl shadow-2xl p-8 space-y-6">
				<div class="space-y-4">
					<div>
                        <label class="block text-gray-300 text-sm font-medium mb-2">Enter URL</label>
                        <div class="relative">
                            <input
                                v-model="url"
                                type="url"
                                placeholder="https://example.com"
                                class="w-full px-4 py-3 input-gradient border border-gray-700 rounded-lg text-gray-200 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                required
                            >
                            <div class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                                <svg class="h-5 w-5 text-gray-400" v-show="url" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                                </svg>
                            </div>
                        </div>
                    </div>
				</div>

                    <button
                        @click="shortenUrl"
                        class="w-full icon-button text-white font-medium py-3 px-4 rounded-lg flex items-center justify-center space-x-2 group"
                    >
                        <svg class="w-5 h-5 transform group-hover:rotate-12 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                        </svg>
                        <span>Generate Short URL</span>
                    </button>
                </div>

                <div v-if="shortUrl" class="mt-6 p-4 input-gradient rounded-lg border border-gray-700">
                    <p class="text-gray-300 mb-2 flex items-center">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
                        </svg>
                        Your shortened URL:
                    </p>
                    <a :href="shortUrl" target="_blank" class="text-blue-400 hover:text-blue-300 break-all block mb-2">[[ shortUrl ]]</a>
                    <a :href="'/stats/' + shortUrl.split('/').pop()" target="_blank" class="text-sm text-gray-400 hover:text-gray-300 flex items-center">
                        <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                        </svg>
                        View Statistics
                    </a>
                </div>
            </div>

			<footer class="mt-8 text-center text-gray-400 text-sm">
			  <p>
			    Built with ❤️ by
			    <a href="https://pipeops.io" class="text-blue-400 hover:text-blue-300 break-all mb-2">
			      PipeOps
			    </a>
			    using Go and Vue.js
			  </p>
			  <a href="https://github.com/PipeOpsHQ/url-shortner" target="_blank" class="github-link flex justify-center space-x-2 mx-auto px-4 py-2 rounded-lg bg-gray-800/50 backdrop-blur-sm">
			    <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
			      <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd" />
			    </svg>
			    <span>Contribute on GitHub</span>
			  </a>
			</footer>

        </div>
    </div>
<script>
        new Vue({
            el: '#app',
            delimiters: ['[[', ']]'],
            data: {
                url: '',
                shortUrl: ''
            },
            mounted() {
                this.createParticles();
            },
            methods: {
                async shortenUrl() {
                    if (!this.url) return;

                    try {
                        const response = await fetch('/shorten', {
                            method: 'POST',
                            headers: { 'Content-Type': 'application/json' },
                            body: JSON.stringify({ url: this.url })
                        });
                        if (!response.ok) {
                            this.shortUrl = 'Error shortening URL';
                            return;
                        }
                        const data = await response.json();
                        this.shortUrl = data.short_url;
                        this.createSuccessParticles();
                    } catch (error) {
                        console.error(error);
                        this.shortUrl = 'Error shortening URL';
                    }
                },
                createParticles() {
                    const particles = 20;
                    for (let i = 0; i < particles; i++) {
                        setTimeout(() => {
                            const particle = document.createElement('div');
                            particle.className = 'particle';

                            const x = Math.random() * window.innerWidth;
                            const y = Math.random() * window.innerHeight;
                            const size = Math.random() * 3 + 1;
                            const tx = (Math.random() - 0.5) * 200;
                            const ty = (Math.random() - 0.5) * 200;

                            particle.style.cssText = 'left: ' + x + 'px;' +
                                'top: ' + y + 'px;' +
                                'width: ' + size + 'px;' +
                                'height: ' + size + 'px;' +
                                '--tx: ' + tx + 'px;' +
                                '--ty: ' + ty + 'px;' +
                                'animation: particle-animation 3s ease-in infinite;';

                            document.body.appendChild(particle);

                            setTimeout(() => {
                                document.body.removeChild(particle);
                            }, 3000);
                        }, i * 200);
                    }
                },
                createSuccessParticles() {
                    this.createParticles();
                }
            }
        });
    </script>
</body>
</html>
`))

// Stats page template with dark theme
// Stats page template with animated gradients
var statsTemplate = template.Must(template.New("stats").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>URL Statistics</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
    <style>
        @keyframes gradientBG {
            0% { background-position: 0% 50%; }
            50% { background-position: 100% 50%; }
            100% { background-position: 0% 50%; }
        }

        @keyframes float {
            0% { transform: translateY(0px); }
            50% { transform: translateY(-10px); }
            100% { transform: translateY(0px); }
        }

        body {
            background: linear-gradient(-45deg, #1a237e, #121836, #2a3f9d, #1e3a8a);
            background-size: 400% 400%;
            animation: gradientBG 15s ease infinite;
            min-height: 100vh;
        }

        .card-gradient {
            background: linear-gradient(180deg, rgba(30, 41, 59, 0.9) 0%, rgba(15, 23, 42, 0.9) 100%);
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255, 255, 255, 0.1);
        }

        .floating-card {
            animation: float 6s ease-in-out infinite;
        }

        .stat-card {
            background: linear-gradient(45deg, rgba(30, 41, 59, 0.8), rgba(15, 23, 42, 0.8));
            backdrop-filter: blur(5px);
        }
    </style>
</head>
<body>
    <div class="min-h-screen flex items-center justify-center p-4">
        <div class="max-w-2xl w-full floating-card">
            <div class="card-gradient rounded-xl shadow-2xl p-8">
                <h1 class="text-3xl font-bold mb-6 text-blue-200 flex items-center">
                    <svg class="w-8 h-8 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                    </svg>
                    URL Statistics
                </h1>

                <div class="space-y-6">
                    <div class="stat-card p-4 rounded-lg border border-gray-700">
                        <p class="text-gray-400 text-sm mb-1">Original URL</p>
                        <a href="{{.LongURL}}" target="_blank" class="text-blue-400 hover:text-blue-300 break-all">{{.LongURL}}</a>
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div class="stat-card p-4 rounded-lg border border-gray-700">
                            <p class="text-gray-400 text-sm mb-1">Total Views</p>
                            <p class="text-2xl font-bold text-white">{{.ViewCount}}</p>
                        </div>

                        <div class="stat-card p-4 rounded-lg border border-gray-700">
                            <p class="text-gray-400 text-sm mb-1">Unique Views</p>
                            <p class="text-2xl font-bold text-white">{{.UniqueViewCount}}</p>
                        </div>
                    </div>
                </div>

                <div class="mt-8 text-center">
                    <a href="/" class="text-blue-400 hover:text-blue-300 flex items-center justify-center">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                        </svg>
                        Back to URL Shortener
                    </a>
                </div>
            </div>
        </div>
    </div>
</body>
</html>
`))

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

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
