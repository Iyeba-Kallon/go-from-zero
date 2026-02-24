package main

import "fmt"

func main() {

    // --- Declaration styles ---

    // Short declaration (most common)
    name := "John"
    age := 20

    // Explicit declaration
    var country string = "Ghana"
    var year int = 2024

    // Declare without value (gets zero value)
    var score int      // defaults to 0
    var isReady bool   // defaults to false
    var title string   // defaults to ""

    // Multiple variables at once
    x, y, z := 1, 2, 3

    // Constants (can never change)
    const Pi = 3.14159
    const AppName = "MyApp"

    // --- Data types ---
    var i int = 10          // integer
    var f float64 = 3.14    // decimal
    var b bool = true       // boolean
    var s string = "hello"  // string

    // --- Print them all ---
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Country:", country)
    fmt.Println("Year:", year)
    fmt.Println("Score:", score)
    fmt.Println("IsReady:", isReady)
    fmt.Println("Title:", title)
    fmt.Println("x, y, z:", x, y, z)
    fmt.Println("Pi:", Pi)
    fmt.Println("AppName:", AppName)
    fmt.Println("int:", i, "float:", f, "bool:", b, "string:", s)

    // --- Type conversion ---
    var myInt int = 42
    var myFloat float64 = float64(myInt)  // int to float
    var backToInt int = int(myFloat)       // float to int
    fmt.Println("Converted:", myInt, myFloat, backToInt)
}