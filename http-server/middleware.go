package main

import (
	"log"
	"net/http"
	"time"
)

// ============================================================
// MIDDLEWARE
//
// Middleware is a function that wraps an http.Handler.
// It runs code BEFORE and/or AFTER the next handler.
//
// The pattern (always the same shape):
//
//   func MyMiddleware(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // --- before ---
//           next.ServeHTTP(w, r)
//           // --- after ---
//       })
//   }
//
// Chaining:  Logger(CORS(router))
//            Requests flow in →→→ Logger → CORS → router
//            Responses flow back ←←← Logger ← CORS ← router
// ============================================================

// ── responseWriter wrapper ────────────────────────────────────
//
// http.ResponseWriter doesn't expose the status code after you've
// written it. We wrap it so Logger can capture it.

type responseWriter struct {
	http.ResponseWriter      // embed the original
	statusCode          int  // captured status
	written             bool // track whether WriteHeader was called
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

// WriteHeader intercepts the status code before passing it through.
func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

// ── Logger middleware ─────────────────────────────────────────

// Logger logs the HTTP method, path, response status, and how long
// the handler took. Runs around every request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // capture start time

		// Wrap the response writer so we can read the status code
		wrapped := newResponseWriter(w)

		next.ServeHTTP(wrapped, r) // call the real handler

		// After the handler returns, log the result
		duration := time.Since(start)
		log.Printf(
			"%-6s %-20s %d  %v",
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			duration,
		)
	})
}

// ── CORS middleware ───────────────────────────────────────────

// CORS adds permissive Cross-Origin Resource Sharing headers so
// the static/index.html page can call this API from the browser.
//
// In production you'd restrict AllowedOrigins to your real domain.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Preflight request: browser sends OPTIONS before the real request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent) // 204
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ── Chain helper ──────────────────────────────────────────────

// Chain applies a list of middleware functions to a handler,
// wrapping them from outermost to innermost.
//
//	Chain(h, Logger, CORS)
//	→ Logger(CORS(h))   — Logger runs first on the way in
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	// Apply in reverse so the first middleware in the list is outermost
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
