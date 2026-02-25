package main

import "fmt"

// ============================================================
// POINTERS IN GO
//
// A pointer holds the MEMORY ADDRESS of a value.
// Instead of copying a value, you share a reference to it.
// This matters for:
//   - Mutating values inside functions
//   - Avoiding expensive copies of large structs
//   - Building linked structures (trees, lists)
// ============================================================

func main() {

	// --- 1. What is a pointer? ---
	fmt.Println("=== Basics ===")

	x := 42
	p := &x // & gives you the address of x

	fmt.Println("x =", x)   // value: 42
	fmt.Println("p =", p)   // address: 0xc000...
	fmt.Println("*p =", *p) // dereference: 42 (value at that address)

	// Modify through pointer
	*p = 100
	fmt.Println("After *p = 100, x =", x) // x changed too!

	// --- 2. Why do we need pointers? ---
	// Without pointers, Go passes COPIES of values to functions.
	fmt.Println("\n=== Without pointer (copy) ===")
	a := 5
	noChange(a)
	fmt.Println("a after noChange:", a) // still 5

	fmt.Println("\n=== With pointer (reference) ===")
	withChange(&a)
	fmt.Println("a after withChange:", a) // now 10

	// --- 3. Pointer to a struct ---
	fmt.Println("\n=== Pointer to struct ===")
	type Point struct {
		X, Y int
	}

	pt := Point{X: 3, Y: 4}
	pp := &pt

	// Go auto-dereferences struct pointers — no need for (*pp).X
	pp.X = 99                   // same as (*pp).X = 99
	fmt.Println("pt.X =", pt.X) // 99

	// --- 4. new() — allocates a pointer to zero value ---
	fmt.Println("\n=== new() ===")
	count := new(int) // *int pointing to 0
	fmt.Println("*count =", *count)
	*count = 7
	fmt.Println("*count after set =", *count)

	// --- 5. Nil pointer ---
	// A pointer with no address assigned — like null in other languages
	fmt.Println("\n=== Nil pointer ===")
	var np *int // nil by default
	fmt.Println("np is nil:", np == nil)
	// NEVER dereference a nil pointer — it will panic!
	// fmt.Println(*np) // ← would crash

	// --- 6. When to use pointers ---
	// ✅ When you want a function to modify a value
	// ✅ When you have a large struct and don't want to copy it
	// ✅ When you need to represent "no value" (nil)
	// ❌ Don't use for basic types like int, string — just pass them directly

	// ============================================================
	// QUICK REFERENCE:
	//
	//  &x          → get address of x (pointer)
	//  *p          → dereference pointer (get value)
	//  func f(p *int) { *p = 5 }   → mutate via pointer
	//  new(T)      → allocate zero value of T, returns *T
	//  var p *int  → nil pointer
	// ============================================================
}

// Takes a copy — original is unchanged
func noChange(n int) {
	n = 999
}

// Takes a pointer — modifies the original
func withChange(n *int) {
	*n = 10
}
