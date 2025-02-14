package main

import "html/template"

// Home page template with animated gradients and iconic buttons
// Home page template with enhanced design
var homeTemplate = template.Must(template.New("home").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>URL Shortener - Shorten your links instantly</title>
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

			<div class="mt-8 text-center space-y-4">
			    <a href="/history" class="text-blue-400 hover:text-blue-300 flex items-center justify-center">
			        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
			            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
			        </svg>
			        View Your URL History
			    </a>
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

// Add new template for history page
var historyTemplate = template.Must(template.New("history").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Your URL History - URL Shortener</title>
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
    <div class="min-h-screen p-6">
        <div class="max-w-6xl mx-auto">
            <div class="floating-header flex flex-col md:flex-row items-center justify-between mb-8 gap-4">
                <div class="text-center md:text-left">
                    <h1 class="text-4xl font-bold text-blue-200 glow-text mb-2">URL History</h1>
                    <p class="text-gray-400">Track and manage your shortened URLs</p>
                </div>
                <a href="/" class="text-blue-400 hover:text-blue-300 flex items-center px-4 py-2 rounded-lg bg-gray-800/50 backdrop-blur-sm transition-all hover:bg-gray-800/70">
                    <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                    </svg>
                    Back to Shortener
                </a>
            </div>

            <div class="card-gradient rounded-xl p-6">
                {{if .}}
                    <div class="grid gap-4">
                        {{range .}}
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
                                            <a href="/{{.ShortCode}}" class="text-blue-400 hover:text-blue-300 break-all text-sm bg-gray-800/50 px-3 py-1 rounded-md hover:bg-gray-800/70 transition-all">
                                                {{.ShortCode}}
                                            </a>
                                            <div class="flex gap-2">
                                                <span class="badge px-2 py-1 rounded-md text-xs text-blue-300">
                                                    {{.UserInfo.Browser}}
                                                </span>
                                                <span class="badge px-2 py-1 rounded-md text-xs text-blue-300">
                                                    {{.UserInfo.OS}}
                                                </span>
                                                <span class="badge px-2 py-1 rounded-md text-xs text-blue-300">
                                                    {{.UserInfo.Device}}
                                                </span>
                                            </div>
                                        </div>

                                        <div class="flex items-center justify-between mt-2">
                                            <p class="text-xs text-gray-400">
                                                Created: {{.CreatedAt.Format "15:04:05"}}
                                            </p>
                                            <a href="/stats/{{.ShortCode}}"
                                               class="text-blue-400 hover:text-blue-300 flex items-center gap-2 text-sm bg-gray-800/50 px-3 py-1 rounded-md hover:bg-gray-800/70 transition-all">
                                                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                                          d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                                                </svg>
                                                View Stats
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
</body>
</html>
`))
