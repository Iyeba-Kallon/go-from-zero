package main

import (
	"fmt"
	"math"
)

// ============================================================
// INTERFACES IN GO
//
// An interface defines a SET OF METHODS a type must have.
// Any type that has those methods automatically satisfies
// the interface — no "implements" keyword needed.
//
// This is how Go achieves flexible, decoupled design.
// ============================================================

// --- 1. Define an interface ---
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- 2. Types that satisfy Shape ---

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// --- 3. A function that accepts ANY Shape ---
// This is the power of interfaces: write once, works for all types
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f | Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ============================================================
// THE Stringer INTERFACE (from fmt package)
// If your type has String() string, fmt uses it automatically
// ============================================================
type Color int

const (
	Red   Color = iota // 0
	Green              // 1
	Blue               // 2
)

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Unknown"
	}
}

// ============================================================
// INTERFACE COMPOSITION
// Interfaces can embed other interfaces
// ============================================================
type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}

type ReadWriter interface {
	Reader // embeds Reader
	Writer // embeds Writer
}

// ============================================================
// THE EMPTY INTERFACE: any
// Accepts any type — used for unknown/mixed data
// (like interface{} in older Go code)
// ============================================================
func describe(v any) {
	fmt.Printf("value: %v | type: %T\n", v, v)
}

// ============================================================
// TYPE ASSERTION & TYPE SWITCH
// Reverse of interface — extract the concrete type
// ============================================================
func whatIsIt(v any) {
	switch val := v.(type) {
	case int:
		fmt.Printf("int: %d\n", val)
	case string:
		fmt.Printf("string: %q\n", val)
	case bool:
		fmt.Printf("bool: %v\n", val)
	case Shape:
		fmt.Printf("Shape with area: %.2f\n", val.Area())
	default:
		fmt.Printf("unknown type: %T\n", val)
	}
}

func main() {

	// --- Shapes ---
	fmt.Println("=== Interface polymorphism ===")
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Print("Circle → ")
	printShapeInfo(c)
	fmt.Print("Rectangle → ")
	printShapeInfo(r)

	// --- Store different types in a slice of interfaces ---
	fmt.Println("\n=== Slice of interfaces ===")
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 2, Height: 5},
		Circle{Radius: 7},
	}
	for _, s := range shapes {
		printShapeInfo(s)
	}

	// --- Stringer ---
	fmt.Println("\n=== Stringer ===")
	fmt.Println(Red, Green, Blue)

	// --- Empty interface / any ---
	fmt.Println("\n=== any ===")
	describe(42)
	describe("hello")
	describe(true)
	describe(Circle{Radius: 2})

	// --- Type switch ---
	fmt.Println("\n=== Type switch ===")
	whatIsIt(10)
	whatIsIt("Go")
	whatIsIt(Circle{Radius: 1})

	// --- Type assertion (single type) ---
	fmt.Println("\n=== Type assertion ===")
	var s Shape = Circle{Radius: 4}
	if circle, ok := s.(Circle); ok {
		fmt.Println("It's a circle! Radius:", circle.Radius)
	}

	// ============================================================
	// QUICK REFERENCE:
	//
	//  type MyInterface interface { Method() ReturnType }
	//  // Any type with Method() satisfies MyInterface — automatically
	//  func DoSomething(v MyInterface) { v.Method() }
	//  val, ok := v.(ConcreteType) // type assertion
	//  switch v.(type) { case int: ... } // type switch
	//  any = accepts any value (replaces interface{})
	// ============================================================
}
