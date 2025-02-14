package main

import (
    "bytes"
    "html/template"
    "log"
    "net/http"
)

// HistoryData represents the template data structure
type HistoryData struct {
    URLs   []URLCreation
    Domain string
}

// HandleHistory handles the URL history page request
func (us *URLShortener) HandleHistory(w http.ResponseWriter, r *http.Request) {
    ip := getIP(r)

    us.mu.RLock()
    urls := us.userHistory[ip]
    us.mu.RUnlock()

    data := HistoryData{
        URLs:   urls,
        Domain: r.Host,
    }

    // Create a buffer to hold the template output
    buf := &bytes.Buffer{}

    // Execute template into buffer first
    if err := historyTemplate.Execute(buf, data); err != nil {
        log.Printf("Template execution error: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set headers
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    // Write the buffered content
    if _, err := buf.WriteTo(w); err != nil {
        log.Printf("Error writing response: %v", err)
    }
}

var historyTemplate = template.Must(template.New("history").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Your URL History - URL Shortener</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- SEO Meta Tags -->
    <meta name="description" content="Shorten your URLs easily and quickly with our URL shortener service. Share your links without the clutter of long URLs.">
    <meta name="keywords" content="URL shortener, link shortener, shorten URL, free URL shortener">
    <meta name="robots" content="index, follow">
	<link rel="canonical" href="https://url.pipeops.app/">

	<!-- Favicon and Apple Touch Icon -->
	<link rel="icon" type="image/png" href="https://pipeops.io/apple-touch-icon.png">
	<link rel="apple-touch-icon" href="https://pipeops.io/apple-touch-icon.png">

	<!-- Open Graph / Social Sharing Image -->
	<meta property="og:image" content="https://pipeops.io/apple-touch-icon.png">

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

        @keyframes glow {
            0%, 100% { box-shadow: 0 0 5px rgba(59, 130, 246, 0.5); }
            50% { box-shadow: 0 0 20px rgba(59, 130, 246, 0.8); }
        }

        @keyframes slideIn {
            from { transform: translateY(20px); opacity: 0; }
            to { transform: translateY(0); opacity: 1; }
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

        .url-card {
            transition: all 0.3s ease;
            animation: slideIn 0.5s ease-out forwards;
            opacity: 0;
        }

        .url-card:hover {
            transform: translateY(-2px) scale(1.01);
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
            background: linear-gradient(45deg, rgba(30, 41, 59, 0.9), rgba(15, 23, 42, 0.95));
        }

        .badge {
            background: linear-gradient(45deg, rgba(59, 130, 246, 0.2), rgba(37, 99, 235, 0.2));
            border: 1px solid rgba(59, 130, 246, 0.3);
            transition: all 0.3s ease;
        }

        .badge:hover {
            background: linear-gradient(45deg, rgba(59, 130, 246, 0.3), rgba(37, 99, 235, 0.3));
            transform: translateY(-1px);
        }

        .glow-text {
            text-shadow: 0 0 10px rgba(59, 130, 246, 0.5);
        }

        .floating-header {
            animation: float 6s ease-in-out infinite;
        }

        .url-card:nth-child(1) { animation-delay: 0.1s; }
        .url-card:nth-child(2) { animation-delay: 0.2s; }
        .url-card:nth-child(3) { animation-delay: 0.3s; }
        .url-card:nth-child(4) { animation-delay: 0.4s; }
        .url-card:nth-child(5) { animation-delay: 0.5s; }
    </style>
</head>
<body class="text-gray-100">
    <div class="min-h-screen p-6" id="history-app">
        <div class="max-w-6xl mx-auto">
            <div class="floating-header flex flex-col md:flex-row items-center justify-between mb-8 gap-4">
                <div class="text-center md:text-left">
                    <h1 class="text-4xl font-bold text-blue-200 glow-text mb-2">URL History</h1>
                    <p class="text-gray-400">Track and manage your shortened URLs</p>
                </div>
                <div class="flex items-center gap-4">
                    <button @click="showHelp = true" class="text-blue-400 hover:text-blue-300 flex items-center px-4 py-2 rounded-lg bg-gray-800/50 backdrop-blur-sm transition-all hover:bg-gray-800/70">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                        Help
                    </button>
                    <a href="/" class="text-blue-400 hover:text-blue-300 flex items-center px-4 py-2 rounded-lg bg-gray-800/50 backdrop-blur-sm transition-all hover:bg-gray-800/70">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                        </svg>
                        New URL
                    </a>
                </div>
            </div>

            <!-- Help Modal -->
            <div v-if="showHelp" class="fixed inset-0 flex items-center justify-center z-50 bg-black/50 backdrop-blur-sm">
                <div class="card-gradient rounded-xl p-8 max-w-lg w-full mx-4 shadow-2xl">
                    <h2 class="text-2xl font-bold text-blue-200 mb-4">How to Use</h2>
                    <div class="space-y-4 text-gray-300">
                        <p>• <strong>Copy URL:</strong> Click the copy icon next to any shortened URL to copy it to your clipboard.</p>
                        <p>• <strong>Share URL:</strong> Use the share icon to quickly share your shortened URL on supported platforms.</p>
                        <p>• <strong>View Stats:</strong> Click the "Stats" badge to see detailed analytics for each URL.</p>
                        <p>• <strong>Click Through:</strong> Click the shortened URL directly to visit the original website.</p>
                    </div>
                    <button @click="showHelp = false" class="mt-6 w-full text-blue-400 hover:text-blue-300 py-2 rounded-lg bg-gray-800/50 transition-all hover:bg-gray-800/70">
                        Got it
                    </button>
                </div>
            </div>

           <div class="card-gradient rounded-xl p-6">
            {{if gt (len .URLs) 0}}
                <div class="grid gap-4">
                    {{range .URLs}}
                        <div class="url-card rounded-lg p-6 border border-gray-700/50">
                            <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
                                <div class="flex-1 space-y-3">
                                    <div class="flex items-start justify-between">
                                        <h3 class="text-lg font-medium text-blue-300 break-all">{{.LongURL}}</h3>
                                        <span class="text-xs text-gray-500 whitespace-nowrap ml-4">
                                            {{.CreatedAt.Format "Jan 02, 2006"}}
                                        </span>
                                    </div>

                                    <div class="flex flex-wrap items-center gap-2">
                                        <div class="flex items-center gap-2 bg-gray-800/50 px-3 py-1 rounded-md">
                                            <a href="/{{.ShortCode}}" target="_blank" class="text-blue-400 hover:text-blue-300 break-all text-sm">
                                                {{$.Domain}}/{{.ShortCode}}
                                            </a>
                                            <button
                                                @click="copyToClipboard(window.location.protocol + '//' + '{{$.Domain}}/{{.ShortCode}}')"
                                                class="text-blue-400 hover:text-blue-300 transition-colors"
                                                :class="{ 'text-green-400': copySuccess['{{.ShortCode}}'] }"
                                                title="Copy shortened URL"
                                            >
                                                <svg v-if="!copySuccess['{{.ShortCode}}']" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                                                </svg>
                                                <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                                                </svg>
                                            </button>
                                            <button
                                                @click="shareUrl(window.location.protocol + '//' + '{{$.Domain}}/{{.ShortCode}}')"
                                                class="text-blue-400 hover:text-blue-300 transition-colors"
                                                title="Share URL"
                                            >
                                                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z" />
                                                </svg>
                                            </button>
                                        </div>
                                        <!-- Only show analytics if viewing is being tracked -->
                                        {{if gt .ViewCount 0}}
                                            <div class="badge px-3 py-1 rounded-md text-sm text-blue-300">
                                                <span class="mr-1">👁</span>
                                                {{.ViewCount}} {{if eq .ViewCount 1}}view{{else}}views{{end}}
                                            </div>
                                        {{end}}
                                        {{if gt .UniqueViewCount 0}}
                                            <div class="badge px-3 py-1 rounded-md text-sm text-blue-300">
                                                <span class="mr-1">👤</span>
                                                {{.UniqueViewCount}} {{if eq .UniqueViewCount 1}}unique view{{else}}unique views{{end}}
                                            </div>
                                        {{end}}
                                        <a
                                            href="/stats/{{.ShortCode}}"
                                            class="badge px-3 py-1 rounded-md text-sm text-blue-300 hover:text-blue-200 transition-colors"
                                            title="View detailed statistics"
                                        >
                                            <span class="mr-1">📊</span>
                                            Stats
                                        </a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    {{end}}
                </div>
            {{else}}
                    <div class="text-center py-16">
                        <div class="w-24 h-24 mx-auto mb-6 rounded-full bg-gray-800/50 flex items-center justify-center">
                            <svg class="w-12 h-12 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                      d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                            </svg>
                        </div>
                        <h2 class="text-2xl font-bold text-gray-400 mb-2">No URLs shortened yet</h2>
                        <p class="text-gray-500 mb-6">Start shortening URLs to see your history here</p>
                        <a href="/" class="inline-flex items-center px-4 py-2 text-blue-400 hover:text-blue-300 bg-gray-800/50 rounded-lg hover:bg-gray-800/70 transition-all">
                            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                            </svg>
                            Create your first short URL
                        </a>
                    </div>
                {{end}}
            </div>
        </div>
    </div>

    <script>
    new Vue({
        el: '#history-app',
        data: {
            copySuccess: {},
            showHelp: false
        },
        methods: {
            async copyToClipboard(text) {
                try {
                    await navigator.clipboard.writeText(text);
                    this.$set(this.copySuccess, text.split('/').pop(), true);
                    setTimeout(() => {
                        this.$set(this.copySuccess, text.split('/').pop(), false);
                    }, 2000);
                } catch (err) {
                    console.error('Failed to copy:', err);
                }
            },
            async shareUrl(url) {
                if (navigator.share) {
                    try {
                        await navigator.share({
                            title: 'Shortened URL',
                            text: 'Check out this shortened URL!',
                            url: url
                        });
                    } catch (err) {
                        if (err.name !== 'AbortError') {
                            console.error('Error sharing:', err);
                        }
                    }
                } else {
                    await this.copyToClipboard(url);
                }
            }
        }
    });
    </script>
</body>
</html>
`))
