# URL Shortener

A lightweight URL shortener built with Go for the backend and Vue.js via CDN for the frontend. This project provides a
simple service to shorten long URLs into compact, shareable links.

## Overview

This URL shortener allows users to:

- Convert long URLs into short, easy-to-share links.
- Redirect from the short URL to the original long URL.
- Enjoy a responsive, modern user interface powered by Vue.js and styled with Tailwind CSS (both loaded via CDN).

## Features

- **URL Shortening:** Generate concise, shareable URLs.
- **Redirection:** Automatically redirect short URLs to the original long URL.
- **SEO Optimized:** Includes meta tags for improved search engine indexing and social media previews.
- **Responsive Design:** Built with Tailwind CSS for a modern, responsive UI.
- **CDN-Powered Frontend:** Leverages Vue.js and Tailwind CSS via CDN for rapid development and simplicity.

## Tech Stack

- **Backend:** Go (Golang)
- **Frontend:** Vue.js (loaded via CDN)
- **CSS Framework:** Tailwind CSS (loaded via CDN)
- **Templating:** Go's built-in templating engine

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher recommended)
- (Optional) A database setup if you plan to persist shortened URLs (e.g., SQLite, Redis, etc.)

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/PipeOpsHQ/url-shortner.git
   cd url-shortner
