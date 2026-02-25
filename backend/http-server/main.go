package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ============================================================
// HTTP SERVER IN GO
//
// Go's standard library includes a full, production-capable
// HTTP server — no framework needed for the basics.
//
// Key pieces:
//   http.HandleFunc(path, handler) → register a route
//   http.ListenAndServe(addr, nil) → start the server
//   w http.ResponseWriter          → write the response
//   r *http.Request                → read the request
// ============================================================

// --- Data model ---
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// --- In-memory "database" ---
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// ============================================================
// HANDLER FUNCTIONS
// A handler has the signature: func(w http.ResponseWriter, r *http.Request)
// ============================================================

// Helper: write a JSON response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Go API!")
}

// GET /health → health check (common in production)
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// GET /users → list all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    users,
	})
}

// POST /users → create a user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
		return
	}

	// Assign an ID and save
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	writeJSON(w, http.StatusCreated, map[string]any{
		"success": true,
		"data":    newUser,
	})
}

// GET/POST /users/manage → router by method
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsersHandler(w, r)
	case http.MethodPost:
		createUserHandler(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
	}
}

// ============================================================
// MAIN — wire up routes and start server
// ============================================================
func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/users", usersHandler)

	// Server config
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("\nTry these routes:")
	fmt.Println("  GET  http://localhost:8080/")
	fmt.Println("  GET  http://localhost:8080/health")
	fmt.Println("  GET  http://localhost:8080/users")
	fmt.Println("  POST http://localhost:8080/users")
	fmt.Println(`     body: {"name":"Carol","email":"carol@example.com"}`)

	// Start server — blocks forever
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}

// ============================================================
// QUICK REFERENCE:
//
//  mux := http.NewServeMux()
//  mux.HandleFunc("/path", handlerFunc)
//
//  handler: func(w http.ResponseWriter, r *http.Request)
//    r.Method                → "GET", "POST", etc.
//    r.URL.Path              → "/users"
//    r.URL.Query().Get("id") → query param ?id=1
//    json.NewDecoder(r.Body).Decode(&v) → parse body
//    w.Header().Set("Content-Type", "application/json")
//    w.WriteHeader(http.StatusOK)
//    json.NewEncoder(w).Encode(v) → write JSON response
//
//  HTTP status codes:
//    http.StatusOK          200
//    http.StatusCreated     201
//    http.StatusBadRequest  400
//    http.StatusNotFound    404
//    http.StatusInternalServerError 500
// ============================================================
