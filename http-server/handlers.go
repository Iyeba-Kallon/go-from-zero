package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ============================================================
// HANDLERS
//
// Each handler has the signature: func(w http.ResponseWriter, r *http.Request)
//
//   w  → where you write your response (status, headers, body)
//   r  → everything about the incoming request
//
// Key patterns shown here:
//   - writeJSON helper (DRY: don't repeat JSON boilerplate)
//   - Reading query params:  r.URL.Query().Get("id")
//   - Decoding a JSON body:  json.NewDecoder(r.Body).Decode(&v)
//   - Input validation and proper error responses
// ============================================================

// writeJSON sets Content-Type, writes the status code, and encodes
// the data as JSON. Used by every handler — keeps things DRY.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ── GET / ────────────────────────────────────────────────────

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Exact match because ServeMux treats "/" as a catch-all.
	// Return 404 for anything that isn't exactly "/".
	if r.URL.Path != "/" {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("route %q not found", r.URL.Path),
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Welcome to my Go HTTP server! 🚀",
		"docs":    "try GET /health  GET /users  POST /users",
	})
}

// ── GET /health ───────────────────────────────────────────────

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"version": "1.0.0",
	})
}

// ── GET /users  +  GET /users?id=N ───────────────────────────

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// If ?id= query param is provided, look up a single user
	idParam := r.URL.Query().Get("id")
	if idParam != "" {
		id, err := strconv.Atoi(idParam) // convert string → int
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "id must be a number",
			})
			return
		}

		user, found := getUserByID(id)
		if !found {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": fmt.Sprintf("user with id %d not found", id),
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    user,
		})
		return
	}

	// No id param → return the full list
	writeJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"count":   len(getUsers()),
		"data":    getUsers(),
	})
}

// ── POST /users ───────────────────────────────────────────────

// createUserRequest is the expected JSON body shape.
// Using a dedicated struct (not User) means the caller can't
// submit an ID — we always assign IDs ourselves.
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var body createUserRequest

	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid JSON body",
		})
		return
	}

	// --- Validation ---
	if body.Name == "" || body.Email == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "name and email are required",
		})
		return
	}

	// Save and respond
	user := addUser(body.Name, body.Email)
	writeJSON(w, http.StatusCreated, map[string]any{
		"success": true,
		"data":    user,
	})
}
