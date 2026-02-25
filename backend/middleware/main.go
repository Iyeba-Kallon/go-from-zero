package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ============================================================
// MIDDLEWARE IN GO
//
// Middleware is a function that WRAPS around HTTP handlers.
// It runs code BEFORE and/or AFTER the real handler.
//
// Common uses:
//   - Logging requests
//   - Authentication / Authorization
//   - CORS headers
//   - Rate limiting
//   - Request timing
//
// The pattern:
//   func MiddlewareName(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // do something BEFORE
//           next.ServeHTTP(w, r)
//           // do something AFTER
//       })
//   }
// ============================================================

// --- 1. Logging Middleware ---
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass to next handler
		next.ServeHTTP(w, r)

		// Log after request completes
		log.Printf("[%s] %s %s (took %v)",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// --- 2. CORS Middleware ---
// Allows cross-origin requests from browsers (needed for frontend apps)
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// --- 3. Auth Middleware ---
// Checks for a bearer token in the Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "Bearer secret-token" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		// Token is valid, continue
		next.ServeHTTP(w, r)
	})
}

// --- 4. Chain multiple middlewares together ---
// Applies middleware right to left (outermost runs first)
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// ============================================================
// HANDLERS
// ============================================================

func publicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"message": "This is public"}`)
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"message": "This is protected — you're authenticated!"}`)
}

// ============================================================
// MAIN
// ============================================================
func main() {
	mux := http.NewServeMux()

	// Public route — only logging + CORS
	mux.Handle("/public", Chain(
		http.HandlerFunc(publicHandler),
		LoggingMiddleware,
		CORSMiddleware,
	))

	// Protected route — logging + CORS + auth
	mux.Handle("/private", Chain(
		http.HandlerFunc(privateHandler),
		LoggingMiddleware,
		CORSMiddleware,
		AuthMiddleware,
	))

	fmt.Println("Middleware server on http://localhost:8081")
	fmt.Println("\nTest it:")
	fmt.Println("  GET  http://localhost:8081/public")
	fmt.Println("  GET  http://localhost:8081/private          → 401")
	fmt.Println(`  GET  http://localhost:8081/private -H "Authorization: Bearer secret-token"`)

	log.Fatal(http.ListenAndServe(":8081", mux))
}

// ============================================================
// QUICK REFERENCE:
//
//  Middleware signature:
//    func Name(next http.Handler) http.Handler {
//        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//            // before
//            next.ServeHTTP(w, r)
//            // after
//        })
//    }
//
//  Apply single:  mux.Handle("/path", MyMiddleware(handler))
//  Chain many:    Chain(handler, Logger, CORS, Auth)
// ============================================================
