package main

import (
	"errors"
	"fmt"
)

// ============================================================
// ERROR HANDLING IN GO
//
// Go has NO exceptions (try/catch). Instead:
//   - Functions return an error as the LAST return value
//   - Callers check it with:  if err != nil { ... }
//
// This makes errors EXPLICIT and impossible to ignore by accident.
// ============================================================

// --- 1. Returning a basic error ---
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // nil = no error
}

// --- 2. Formatted error with context (preferred) ---
func openFile(filename string) error {
	if filename == "" {
		return fmt.Errorf("openFile: filename cannot be empty")
	}
	// imagine file opening here...
	return nil
}

// ============================================================
// CUSTOM ERROR TYPES
// Define your own error type for richer error info
// Must implement the Error() string method
// ============================================================

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %q: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be positive"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

// ============================================================
// WRAPPING & UNWRAPPING ERRORS
// fmt.Errorf with %w wraps an error so you can check its type
// errors.Is() checks if a specific error is in the chain
// errors.As() extracts a specific error type from the chain
// ============================================================

var ErrNotFound = errors.New("not found")

func findUser(id int) (string, error) {
	users := map[int]string{1: "Alice", 2: "Bob"}
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("findUser(%d): %w", id, ErrNotFound) // wrap
	}
	return name, nil
}

func main() {

	// --- Basic error handling ---
	fmt.Println("=== Basic errors ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// --- fmt.Errorf ---
	fmt.Println("\n=== Formatted errors ===")
	err = openFile("")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = openFile("data.json")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File opened OK")
	}

	// --- Custom error type ---
	fmt.Println("\n=== Custom error type ===")
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Error:", err)
		// Extract the custom type with errors.As
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("  Field: %s\n  Message: %s\n", ve.Field, ve.Message)
		}
	}

	err = validateAge(25)
	if err == nil {
		fmt.Println("Age 25 is valid ✓")
	}

	// --- Wrapped errors (errors.Is) ---
	fmt.Println("\n=== Wrapped errors ===")
	name, err := findUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", name)
	}

	_, err = findUser(99)
	if err != nil {
		fmt.Println("Error:", err)
		// errors.Is checks through the wrapping chain
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ This is a 'not found' error")
		}
	}

	// --- Sentinel errors (for comparing specific errors) ---
	fmt.Println("\n=== Sentinel errors ===")
	fmt.Println("ErrNotFound:", ErrNotFound)

	// ============================================================
	// QUICK REFERENCE:
	//
	//  result, err := someFunc()
	//  if err != nil { handle }               // always check!
	//
	//  errors.New("msg")                       // simple error
	//  fmt.Errorf("context: %w", err)          // wrap with context
	//  errors.Is(err, ErrSomething)            // check type in chain
	//  errors.As(err, &myErrType)              // extract type from chain
	//
	//  Custom: type MyErr struct { ... }
	//          func (e *MyErr) Error() string { ... }
	// ============================================================
}
