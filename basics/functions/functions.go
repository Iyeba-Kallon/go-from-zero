package main

import "fmt"

// ============================================================
// FUNCTIONS IN GO
// Functions are first-class citizens — you can pass them around
// like variables, return them, and store them.
// ============================================================

// --- 1. Basic function ---
func greet(name string) {
	fmt.Println("Hello,", name)
}

// --- 2. Function with a return value ---
func add(a int, b int) int {
	return a + b
}

// --- 3. Multiple return values (very common in Go) ---
// This is how Go handles errors — return the result AND an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil // nil means "no error"
}

// --- 4. Named return values ---
// You can name the return variables and use a bare "return"
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // returns min and max automatically
}

// --- 5. Variadic function (accepts any number of args) ---
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- 6. Functions as values (first-class) ---
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func double(n int) int {
	return n * 2
}

func main() {

	// Call a basic function
	greet("Riya")

	// Call with return value
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// Multiple return values
	fmt.Println("\n--- Multiple returns ---")
	answer, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", answer)
	}

	// Divide by zero — error case
	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Caught error:", err)
	}

	// Named returns
	fmt.Println("\n--- Named returns ---")
	nums := []int{3, 7, 1, 9, 4}
	lo, hi := minMax(nums)
	fmt.Println("Min:", lo, "Max:", hi)

	// Variadic
	fmt.Println("\n--- Variadic ---")
	fmt.Println("Sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("Sum(1..5):", sum(1, 2, 3, 4, 5))
	// You can also spread a slice into a variadic function
	numbers := []int{10, 20, 30}
	fmt.Println("Sum(slice...):", sum(numbers...))

	// Function as a value
	fmt.Println("\n--- Functions as values ---")
	fmt.Println("applyTwice(double, 3):", applyTwice(double, 3)) // 12

	// --- 7. Anonymous function (closure) ---
	// A function defined and used inline, capturing surrounding vars
	fmt.Println("\n--- Closure ---")
	counter := 0
	increment := func() int {
		counter++ // captures the outer variable
		return counter
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3

	// ============================================================
	// QUICK REFERENCE:
	//
	//  func name(param type) returnType { }
	//  func name(a, b int) (int, error) { }   // multiple returns
	//  func name(nums ...int) { }              // variadic
	//  result, err := someFunc()               // unpack returns
	//  if err != nil { handle error }          // always check errors!
	// ============================================================
}
