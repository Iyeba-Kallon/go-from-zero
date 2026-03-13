package main

import "sync"

// ============================================================
// MODELS
//
// This file defines our data types and the in-memory "database".
//
// Key concepts:
//   - struct tags (`json:"name"`) control how fields appear in JSON
//   - sync.Mutex protects shared data when multiple requests
//     come in at the same time (goroutines running concurrently)
// ============================================================

// User is the main data model for this API.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// store holds our in-memory "database" and a mutex to make
// concurrent reads/writes safe.
var store = struct {
	mu     sync.Mutex
	users  []User
	nextID int
}{
	users: []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	},
	nextID: 3, // next user will get ID 3
}

// getUsers returns a copy of all users (safe to read concurrently).
func getUsers() []User {
	store.mu.Lock()
	defer store.mu.Unlock()

	// Return a copy so callers can't accidentally mutate the store
	result := make([]User, len(store.users))
	copy(result, store.users)
	return result
}

// getUserByID finds a user by ID. Returns (user, true) or (zero, false).
func getUserByID(id int) (User, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()

	for _, u := range store.users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

// addUser appends a new user to the store and returns it with its assigned ID.
func addUser(name, email string) User {
	store.mu.Lock()
	defer store.mu.Unlock()

	u := User{
		ID:    store.nextID,
		Name:  name,
		Email: email,
	}
	store.nextID++
	store.users = append(store.users, u)
	return u
}
