package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ============================================================
// MAIN — entry point
//
// This file wires everything together:
//   1. Create the router and register routes
//   2. Serve static files from ./static/
//   3. Wrap the router with middleware
//   4. Configure and start the http.Server
//
// Keeping main.go small and focused is good practice —
// all the logic lives in the other files.
// ============================================================

func main() {
	// ── 1. Router & routes ─────────────────────────────────
	r := NewRouter()

	r.GET("/", homeHandler)
	r.GET("/health", healthHandler)
	r.GET("/users", getUsersHandler)
	r.POST("/users", createUserHandler)

	// ── 2. Static file server ──────────────────────────────
	// http.FileServer serves files from a directory.
	// http.StripPrefix removes "/static" before looking up the file,
	// so a request for /static/index.html reads ./static/index.html.
	fs := http.FileServer(http.Dir("./static"))
	r.mux.Handle("/static/", http.StripPrefix("/static", fs))

	// ── 3. Middleware chain ────────────────────────────────
	// Logger wraps CORS wraps the router.
	// Execution order on a request: Logger → CORS → router → handler
	handler := Chain(r, Logger, CORS)

	// ── 4. Server configuration ────────────────────────────
	// Always set timeouts on production servers to prevent slowloris
	// attacks and resource leaks from idle connections.
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second, // max time to read request
		WriteTimeout: 10 * time.Second, // max time to write response
		IdleTimeout:  60 * time.Second, // max keep-alive idle time
	}

	// Print a friendly startup banner
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║       Go HTTP Server — from scratch      ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  http://localhost:8080                   ║")
	fmt.Println("║                                          ║")
	fmt.Println("║  GET  /                                  ║")
	fmt.Println("║  GET  /health                            ║")
	fmt.Println("║  GET  /users                             ║")
	fmt.Println("║  GET  /users?id=1                        ║")
	fmt.Println("║  POST /users  {name, email}              ║")
	fmt.Println("║  GET  /static/index.html  (demo UI)      ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	// ListenAndServe blocks forever — server runs until you Ctrl+C
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server error:", err)
	}
}
