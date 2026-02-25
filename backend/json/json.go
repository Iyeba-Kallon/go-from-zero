package main

import (
	"encoding/json"
	"fmt"
)

// ============================================================
// JSON IN GO
//
// JSON is the backbone of REST APIs. Go's standard library
// handles it with encoding/json.
//
// Key concepts:
//   json.Marshal   → Go value → JSON bytes
//   json.Unmarshal → JSON bytes → Go value
//   Struct tags    → control field names in JSON
// ============================================================

// --- Struct with JSON tags ---
// Tags tell Go what the JSON key name should be
// `json:"name"`         → use "name" as key
// `json:"name,omitempty"` → skip field if empty/zero
// `json:"-"`            → never include this field
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"` // never expose!
	CreatedAt string `json:"created_at"`
	Bio       string `json:"bio,omitempty"` // skip if empty
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func main() {

	// ============================================================
	// 1. MARSHAL — Go struct → JSON string
	// ============================================================
	fmt.Println("=== Marshal (Go → JSON) ===")

	user := User{
		ID:        1,
		Name:      "Alice",
		Email:     "alice@example.com",
		Password:  "secret123",
		CreatedAt: "2024-01-15",
	}

	// Marshal returns []byte (raw bytes)
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Compact JSON:")
	fmt.Println(string(data))

	// MarshalIndent for pretty-printing
	pretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("\nPretty JSON:")
	fmt.Println(string(pretty))

	// ============================================================
	// 2. UNMARSHAL — JSON string → Go struct
	// ============================================================
	fmt.Println("\n=== Unmarshal (JSON → Go) ===")

	jsonStr := `{
		"id": 2,
		"name": "Bob",
		"email": "bob@example.com",
		"created_at": "2024-02-20",
		"bio": "Backend developer"
	}`

	var u User
	err = json.Unmarshal([]byte(jsonStr), &u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s | Email: %s | Bio: %s\n", u.Name, u.Email, u.Bio)
	fmt.Println("Password (should be empty):", u.Password)

	// ============================================================
	// 3. SLICE OF STRUCTS (JSON array)
	// ============================================================
	fmt.Println("\n=== JSON array ===")

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	arr, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(arr))

	// Parse JSON array
	jsonArr := `[{"id":3,"name":"Carol","email":"carol@example.com","created_at":"2024-03-01"}]`
	var parsed []User
	json.Unmarshal([]byte(jsonArr), &parsed)
	fmt.Println("Parsed:", parsed[0].Name, parsed[0].Email)

	// ============================================================
	// 4. ENCODING TO map[string]any (dynamic/unknown JSON)
	// ============================================================
	fmt.Println("\n=== Dynamic JSON with map ===")

	rawJSON := `{"status": "ok", "count": 42, "tags": ["go", "backend"]}`
	var result map[string]any
	json.Unmarshal([]byte(rawJSON), &result)

	fmt.Println("status:", result["status"])
	fmt.Println("count:", result["count"])
	fmt.Println("tags:", result["tags"])

	// ============================================================
	// 5. API RESPONSE PATTERN (common in backends)
	// ============================================================
	fmt.Println("\n=== API Response pattern ===")

	resp := APIResponse{
		Success: true,
		Message: "User fetched",
		Data:    user,
	}
	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(respJSON))

	// ============================================================
	// QUICK REFERENCE:
	//
	//  json.Marshal(v)              → any → []byte
	//  json.MarshalIndent(v,"","  ") → pretty []byte
	//  json.Unmarshal(data, &v)     → []byte → struct
	//  string(data)                 → []byte → string
	//  []byte(str)                  → string → []byte
	//
	//  Struct tags:
	//    `json:"key"`               → rename field
	//    `json:"key,omitempty"`     → skip if zero value
	//    `json:"-"`                 → always exclude
	// ============================================================
}
