package main

import "html/template"

var statsTemplate = template.Must(template.New("stats").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>URL Statistics</title>
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
    <div class="min-h-screen flex items-center justify-center p-4" id="stats-app">
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
                        <div class="flex justify-between items-start">
                            <div class="flex-grow">
                                <p class="text-gray-400 text-sm mb-1">Original URL</p>
                                <a href="{{.LongURL}}" target="_blank" class="text-blue-400 hover:text-blue-300 break-all">{{.LongURL}}</a>
                            </div>
                            <div class="flex space-x-2 ml-4">
                                <button
                                    @click="copyToClipboard('{{.LongURL}}')"
                                    class="p-2 rounded-lg hover:bg-gray-700 transition-colors duration-200"
                                    :class="{ 'text-green-400': copySuccess, 'text-gray-400': !copySuccess }"
                                    title="Copy URL">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path v-if="!copySuccess" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                                        <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                                    </svg>
                                </button>
                                <button
                                    @click="shareUrl('{{.LongURL}}')"
                                    class="p-2 rounded-lg hover:bg-gray-700 transition-colors duration-200 text-gray-400"
                                    title="Share URL">
                                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z" />
                                    </svg>
                                </button>
                            </div>
                        </div>
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
    <script>
    new Vue({
        el: '#stats-app',
        data: {
            copySuccess: false
        },
        methods: {
            async copyToClipboard(text) {
                try {
                    await navigator.clipboard.writeText(text);
                    this.copySuccess = true;
                    setTimeout(() => {
                        this.copySuccess = false;
                    }, 2000);
                } catch (err) {
                    console.error('Failed to copy:', err);
                }
            },
            async shareUrl(url) {
                if (navigator.share) {
                    try {
                        await navigator.share({
                            title: 'URL Statistics',
                            text: 'Check out these URL statistics!',
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
