package main

import "fmt"

// ============================================================
// STRUCTS IN GO
//
// A struct is Go's way of grouping related data together.
// Think of it like a class, but without inheritance.
// You add behavior to structs using METHODS.
// ============================================================

// --- 1. Defining a struct ---
type User struct {
	Name    string
	Email   string
	Age     int
	IsAdmin bool
}

// --- 2. Constructor function (Go convention: NewTypeName) ---
// Go has no built-in constructors — you write a function instead
func NewUser(name, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}

// --- 3. Method on a struct (value receiver) ---
// The struct is COPIED — safe for reads, no mutation
func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s (%s)", u.Name, u.Email)
}

// --- 4. Method with a pointer receiver ---
// The struct is NOT copied — use this to modify fields
func (u *User) Promote() {
	u.IsAdmin = true
}

func (u *User) Birthday() {
	u.Age++
}

// --- 5. String() method — custom printing ---
// If you define String(), fmt will use it automatically
func (u User) String() string {
	role := "user"
	if u.IsAdmin {
		role = "admin"
	}
	return fmt.Sprintf("User{%s | %s | age %d | %s}", u.Name, u.Email, u.Age, role)
}

// --- 6. Nested struct ---
type Address struct {
	City    string
	Country string
}

type Employee struct {
	User    // embedded struct (like inheritance, but not)
	Address Address
	Salary  float64
}

func main() {

	// --- Create a struct ---
	fmt.Println("=== Creating structs ===")

	// Struct literal
	u1 := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   28,
	}
	fmt.Println(u1)

	// Using the constructor
	u2 := NewUser("Bob", "bob@example.com", 35)
	fmt.Println(u2)

	// --- Access fields ---
	fmt.Println("\n=== Field access ===")
	fmt.Println("Name:", u1.Name)
	fmt.Println("Email:", u1.Email)

	// --- Call methods ---
	fmt.Println("\n=== Methods ===")
	fmt.Println(u1.Greet())

	// Pointer receiver method on value — Go handles this automatically
	u1.Promote() // Go takes the address for you
	u1.Birthday()
	fmt.Println("After promote & birthday:", u1)

	// --- Pointer vs Value ---
	fmt.Println("\n=== Pointer receiver ===")
	u3 := NewUser("Carol", "carol@example.com", 22)
	fmt.Println("Before:", *u3)
	u3.Birthday()
	fmt.Println("After birthday:", *u3)

	// --- Nested / Embedded struct ---
	fmt.Println("\n=== Embedded struct ===")
	emp := Employee{
		User:    User{Name: "Dave", Email: "dave@corp.com", Age: 40},
		Address: Address{City: "Freetown", Country: "Sierra Leone"},
		Salary:  75000,
	}
	// You can access embedded fields directly
	fmt.Println("Employee name:", emp.Name) // from embedded User
	fmt.Println("Employee city:", emp.Address.City)
	fmt.Println(emp.Greet()) // methods are promoted too!

	// --- Anonymous struct (quick one-off) ---
	fmt.Println("\n=== Anonymous struct ===")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)

	// ============================================================
	// QUICK REFERENCE:
	//
	//  type T struct { Field Type }      // define
	//  t := T{Field: value}              // create
	//  t.Field                           // access
	//  func (t T) Method() { }           // value receiver (read)
	//  func (t *T) Method() { }          // pointer receiver (write)
	//  func NewT(...) *T { return &T{} } // constructor pattern
	// ============================================================
}
