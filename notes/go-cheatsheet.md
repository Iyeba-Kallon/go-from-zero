# Go Cheatsheet

## Running Go

```bash
go run filename.go   # run a file
go build filename.go # compile to executable
```

---

## Program Structure

```go
package main        // every executable needs this

import "fmt"        // import packages

func main() {       // entry point, runs first
    // your code here
}
```

---

## Printing

```go
fmt.Println("Hello")                          // prints with newline
fmt.Print("Hello ")                           // prints without newline
fmt.Printf("Name: %s Age: %d\n", name, age)  // formatted
```

### Format Verbs

| Verb | Meaning                        |
|------|--------------------------------|
| `%s` | string                         |
| `%d` | integer                        |
| `%f` | float                          |
| `%b` | boolean                        |
| `%v` | any value (useful for debugging) |
| `\n` | newline                        |

---

## Variables

```go
// Short declaration (most common, inside functions only)
name := "John"
age := 20

// Explicit declaration
var country string = "Ghana"
var year int = 2024

// Declare without value (gets zero value)
var score int      // 0
var isReady bool   // false
var title string   // ""

// Multiple at once
x, y, z := 1, 2, 3

// Constants
const Pi = 3.14159
const AppName = "MyApp"
```

---

## Data Types

```go
var i int     = 10      // whole numbers
var f float64 = 3.14    // decimals
var b bool    = true    // true / false
var s string  = "hello" // text
```

### Zero Values

| Type      | Zero Value |
|-----------|------------|
| `int`     | `0`        |
| `float64` | `0.0`      |
| `bool`    | `false`    |
| `string`  | `""`       |

---

## Type Conversion

```go
var myInt int = 42
var myFloat float64 = float64(myInt)  // int → float
var backToInt int = int(myFloat)       // float → int

// Go never converts types automatically — always explicit
```
