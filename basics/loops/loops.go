package main

import "fmt"

func main() {

	// ============================================================
	// LOOPS IN GO
	// Go only has ONE loop keyword: "for"
	// But it can behave like a while loop, infinite loop, and more.
	// ============================================================

	// --- 1. Classic for loop (like C/Java) ---
	// for init; condition; post { ... }
	fmt.Println("--- Classic for loop ---")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
	// Output: 0, 1, 2, 3, 4

	// --- 2. For as a While loop ---
	// Just use the condition, drop the init and post.
	fmt.Println("\n--- For as a while loop ---")
	count := 1
	for count <= 5 {
		fmt.Println("count =", count)
		count++
	}

	// --- 3. Infinite loop ---
	// Omit everything. Use "break" to escape.
	fmt.Println("\n--- Infinite loop with break ---")
	n := 0
	for {
		if n == 3 {
			break // stops the loop entirely
		}
		fmt.Println("n =", n)
		n++
	}

	// --- 4. continue ---
	// Skips the rest of the current iteration and jumps to the next.
	fmt.Println("\n--- Using continue (skip even numbers) ---")
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("odd:", i)
	}

	// --- 5. for range over a slice (array) ---
	// "range" gives you both the index and the value.
	fmt.Println("\n--- Range over a slice ---")
	fruits := []string{"mango", "banana", "pawpaw"}
	for index, fruit := range fruits {
		fmt.Println(index, "->", fruit)
	}

	// Use _ to ignore the index if you don't need it
	fmt.Println("\n--- Range, ignoring index ---")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// --- 6. for range over a string ---
	// Iterates character by character (as runes).
	fmt.Println("\n--- Range over a string ---")
	word := "Go!"
	for i, ch := range word {
		fmt.Printf("index %d: %c\n", i, ch)
	}

	// --- 7. for range over a map ---
	// Maps have no guaranteed order, but range still works.
	fmt.Println("\n--- Range over a map ---")
	capitals := map[string]string{
		"Ghana":        "Accra",
		"Nigeria":      "Abuja",
		"Sierra Leone": "Freetown",
	}
	for country, capital := range capitals {
		fmt.Println(country, "->", capital)
	}

	// --- 8. Nested loops ---
	// A loop inside a loop. Common for grids/tables.
	fmt.Println("\n--- Nested loops (multiplication table) ---")
	for row := 1; row <= 3; row++ {
		for col := 1; col <= 3; col++ {
			fmt.Printf("%d x %d = %d\t", row, col, row*col)
		}
		fmt.Println() // newline after each row
	}

	// --- 9. Labeled break ---
	// Break out of an OUTER loop from inside a nested loop.
	fmt.Println("\n--- Labeled break ---")
outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i+j == 4 {
				fmt.Printf("Breaking at i=%d, j=%d\n", i, j)
				break outer // exits the outer loop entirely
			}
		}
	}

	// ============================================================
	// QUICK REFERENCE:
	//
	//  Classic:        for i := 0; i < n; i++ { }
	//  While-style:    for condition { }
	//  Infinite:       for { }
	//  Range:          for i, v := range collection { }
	//  Skip index:     for _, v := range collection { }
	//  break           → exits the loop
	//  continue        → skips to the next iteration
	// ============================================================
}
