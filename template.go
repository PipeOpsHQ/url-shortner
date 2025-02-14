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

        .input-gradient {
            background: linear-gradient(to bottom, rgba(30, 41, 59, 0.8), rgba(15, 23, 42, 0.8));
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
                                @input="validateUrl"
                                @paste="handlePaste"
                                class="w-full px-4 py-3 input-gradient border rounded-lg text-gray-200 placeholder-gray-500 focus:outline-none focus:ring-2 focus:border-transparent transition-colors duration-200"
                                :class="{
                                    'border-gray-700': !url || (url && isValidUrl),
                                    'border-red-500 focus:ring-red-500': url && !isValidUrl,
                                    'border-green-500 focus:ring-green-500': url && isValidUrl
                                }"
                                required
                            >
                            <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                                <!-- Valid URL Icon -->
                                <svg v-if="url && isValidUrl" class="h-5 w-5 text-green-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                                </svg>
                                <!-- Invalid URL Icon -->
                                <svg v-else-if="url && !isValidUrl" class="h-5 w-5 text-red-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                                </svg>
                                <!-- URL Icon -->
                                <svg v-else class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                                </svg>
                            </div>
                        </div>
                        <!-- Validation Message -->
                        <transition
                            enter-active-class="transition ease-out duration-200"
                            enter-from-class="opacity-0 -translate-y-1"
                            enter-to-class="opacity-100 translate-y-0"
                            leave-active-class="transition ease-in duration-150"
                            leave-from-class="opacity-100 translate-y-0"
                            leave-to-class="opacity-0 -translate-y-1"
                        >
                            <p v-if="url && !isValidUrl" class="mt-2 text-red-400 text-sm">
                                Please enter a valid URL (e.g., https://example.com)
                            </p>
                        </transition>
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
                    <div class="flex items-center gap-2 mb-3">
                        <a :href="shortUrl" target="_blank" class="text-blue-400 hover:text-blue-300 break-all flex-1">[[ shortUrl ]]</a>
                        <button
                            @click="copyToClipboard(shortUrl)"
                            class="p-2 text-blue-400 hover:text-blue-300 bg-gray-800/50 rounded-lg transition-all hover:bg-gray-800/70"
                            :title="copySuccess ? 'Copied!' : 'Copy to clipboard'"
                        >
                            <svg v-if="!copySuccess" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                            </svg>
                            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                            </svg>
                        </button>
                        <button
                            @click="shareUrl(shortUrl)"
                            class="p-2 text-blue-400 hover:text-blue-300 bg-gray-800/50 rounded-lg transition-all hover:bg-gray-800/70"
                            title="Share URL"
                        >
                            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z" />
                            </svg>
                        </button>
                    </div>
                    <div class="flex items-center justify-between">
                        <a :href="'/stats/' + shortUrl.split('/').pop()" class="text-sm text-gray-400 hover:text-gray-300 flex items-center">
                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                            </svg>
                            View Statistics
                        </a>
                        <span v-if="copySuccess" class="text-sm text-green-400 animate-fade-in">Copied to clipboard!</span>
                    </div>
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
    </div>
<script>
new Vue({
    el: '#app',
    delimiters: ['[[', ']]'],
    data: {
        url: '',
        shortUrl: '',
        copySuccess: false
    },
    mounted() {
        this.createParticles();
    },
    methods: {
        validateUrl() {
            if (!this.url) {
                this.isValidUrl = false;
                return;
            }

            try {
                // Try to construct a URL object
                const urlObj = new URL(this.url);
                // Check if protocol is http or https
                this.isValidUrl = urlObj.protocol === 'http:' || urlObj.protocol === 'https:';

                // Additional validation for minimum domain length
                const hostParts = urlObj.hostname.split('.');
                if (hostParts.length < 2 || hostParts.some(part => part.length === 0)) {
                    this.isValidUrl = false;
                }
            } catch {
                this.isValidUrl = false;
            }
        },
        handlePaste(e) {
            // Allow paste event to complete before validation
            setTimeout(() => {
                this.validateUrl();
            }, 0);
        },
        autoCorrectUrl() {
            if (this.url && !this.url.startsWith('http://') && !this.url.startsWith('https://')) {
                this.url = 'https://' + this.url;
                this.validateUrl();
            }
        },
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
        },
        async shortenUrl() {
            if (this.isValidUrl) {
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
    },
    watch: {
        url(newVal) {
            // Auto-correct URL after a short delay when user stops typing
            clearTimeout(this.urlTimer);
            this.urlTimer = setTimeout(() => {
                this.autoCorrectUrl();
            }, 1000);
        }
    },
    beforeDestroy() {
        clearTimeout(this.urlTimer);
    }
});
</script>
</body>
</html>
`))
