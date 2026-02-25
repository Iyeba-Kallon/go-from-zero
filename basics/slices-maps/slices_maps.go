package main

import "fmt"

// ============================================================
// SLICES & MAPS IN GO
//
// Slice = a dynamic list (like an array that can grow)
// Map   = a dictionary / key-value store (like a hashmap)
// ============================================================

func main() {

	// ===========================
	// SLICES
	// ===========================

	// --- 1. Creating slices ---
	fmt.Println("=== SLICES ===")

	// Slice literal
	fruits := []string{"mango", "banana", "pawpaw"}
	fmt.Println("fruits:", fruits)

	// Empty slice with make(type, length, capacity)
	scores := make([]int, 0, 5) // len=0, cap=5

	// --- 2. append — add items ---
	scores = append(scores, 90, 85, 78)
	fmt.Println("scores:", scores)

	// --- 3. Indexing ---
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Last fruit:", fruits[len(fruits)-1])

	// --- 4. Slicing a slice (sub-slice) ---
	// [start:end] — end is exclusive
	fmt.Println("fruits[0:2]:", fruits[0:2]) // mango, banana
	fmt.Println("fruits[1:]:", fruits[1:])   // banana, pawpaw
	fmt.Println("fruits[:2]:", fruits[:2])   // mango, banana

	// --- 5. Iterating ---
	fmt.Println("\n--- Range over slice ---")
	for i, f := range fruits {
		fmt.Printf("%d: %s\n", i, f)
	}

	// --- 6. 2D slice (slice of slices) ---
	fmt.Println("\n--- 2D slice ---")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, row := range matrix {
		fmt.Println(row)
	}

	// --- 7. Copy a slice (slices share memory!) ---
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 99
	fmt.Println("\noriginal:", original) // unchanged
	fmt.Println("copied:", copied)       // modified copy

	// ===========================
	// MAPS
	// ===========================
	fmt.Println("\n=== MAPS ===")

	// --- 1. Map literal ---
	person := map[string]string{
		"name":    "Iyeba",
		"country": "Sierra Leone",
		"role":    "developer",
	}
	fmt.Println("person:", person)

	// --- 2. Create empty map with make ---
	ages := make(map[string]int)

	// --- 3. Add & update values ---
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Alice"] = 26 // update
	fmt.Println("\nages:", ages)

	// --- 4. Read a value ---
	fmt.Println("Alice's age:", ages["Alice"])

	// --- 5. Check if a key exists ---
	// The comma-ok pattern: value, ok := map[key]
	age, ok := ages["Charlie"]
	if ok {
		fmt.Println("Charlie:", age)
	} else {
		fmt.Println("Charlie not found")
	}

	// --- 6. Delete a key ---
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// --- 7. Iterate over a map ---
	fmt.Println("\n--- Range over map ---")
	capitals := map[string]string{
		"Sierra Leone": "Freetown",
		"Ghana":        "Accra",
		"Nigeria":      "Abuja",
	}
	for country, capital := range capitals {
		fmt.Printf("%s → %s\n", country, capital)
	}

	// ===========================
	// SLICE OF MAPS (common in backend!)
	// ===========================
	fmt.Println("\n=== Slice of Maps (like JSON array) ===")
	users := []map[string]string{
		{"name": "Alice", "role": "admin"},
		{"name": "Bob", "role": "user"},
		{"name": "Carol", "role": "user"},
	}
	for _, u := range users {
		fmt.Printf("Name: %s | Role: %s\n", u["name"], u["role"])
	}

	// ============================================================
	// QUICK REFERENCE:
	//
	//  Slice:
	//    nums := []int{1, 2, 3}
	//    nums = append(nums, 4)
	//    nums[0:2]             → sub-slice
	//    len(nums)             → length
	//    copy(dst, src)        → copy
	//
	//  Map:
	//    m := map[string]int{"a": 1}
	//    m["b"] = 2            → add/update
	//    val, ok := m["key"]   → safe read
	//    delete(m, "key")      → remove
	// ============================================================
}
