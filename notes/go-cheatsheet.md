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
name := "Riya"
age := 21

// Explicit declaration
var country string = "Sierra Leone"
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

---

## Loops

Go only has one loop keyword: `for`.

```go
// 1. Classic for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 2. "While" style loop
count := 1
for count <= 5 {
    fmt.Println(count)
    count++
}

// 3. Infinite loop
for {
    if done {
        break
    }
}

// 4. Range loop (for arrays, slices, maps, strings)
fruits := []string{"apple", "banana"}
for index, value := range fruits {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}

// 5. Loop control
continue // skip rest of current iteration
break    // exit loop entirely
```

---

## Functions

```go
// Basic
func greet(name string) string { return "Hello, " + name }

// Multiple returns (Go's error pattern)
func divide(a, b float64) (float64, error) {
    if b == 0 { return 0, errors.New("divide by zero") }
    return a / b, nil
}
result, err := divide(10, 2)
if err != nil { /* handle */ }

// Variadic
func sum(nums ...int) int { total := 0; for _, n := range nums { total += n }; return total }
sum(1, 2, 3)       // pass individually
sum(slice...)       // spread a slice

// Closures
add := func(a, b int) int { return a + b }
```

---

## Slices

```go
s := []string{"a", "b", "c"}   // literal
s = append(s, "d")             // add element
s[0]                           // index
s[1:3]                         // sub-slice [b c]
len(s)                         // length

dst := make([]string, len(s))
copy(dst, s)                   // deep copy

for i, v := range s { }        // range
```

## Maps

```go
m := map[string]int{"a": 1}   // literal
m["b"] = 2                    // add/update
val := m["a"]                 // read  
val, ok := m["x"]             // safe read (ok=false if missing)
delete(m, "a")                // remove
for k, v := range m { }       // range
```

---

## Pointers

```go
x := 42
p := &x      // pointer to x
*p = 100     // dereference (changes x)

func setTo10(n *int) { *n = 10 }
setTo10(&x)  // pass address

// Use pointer receivers to mutate a struct
func (u *User) Promote() { u.IsAdmin = true }
```

---

## Structs

```go
type User struct {
    Name  string
    Email string
    Age   int
}

// Constructor pattern
func NewUser(name, email string) *User {
    return &User{Name: name, Email: email}
}

// Value receiver (read only)
func (u User) Greet() string { return "Hi " + u.Name }

// Pointer receiver (mutates)
func (u *User) Birthday() { u.Age++ }

u := NewUser("Alice", "alice@x.com")
u.Birthday()
fmt.Println(u.Greet())
```

---

## Interfaces

```go
type Shape interface {
    Area() float64
}

// Any type with Area() satisfies Shape — automatically
type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func printArea(s Shape) { fmt.Println(s.Area()) }

// Type assertion
val, ok := s.(Circle)

// Type switch
switch v := thing.(type) {
case int:    fmt.Println("int:", v)
case string: fmt.Println("string:", v)
}

// any = accepts anything (replaces interface{})
func describe(v any) { fmt.Printf("%T %v\n", v, v) }
```

---

## Errors

```go
// Return & check
result, err := someFunc()
if err != nil { /* handle */ }

// Create errors
errors.New("message")
fmt.Errorf("context: %w", err)   // wraps for chain checking

// Custom error type
type MyError struct{ Message string }
func (e *MyError) Error() string { return e.Message }

// Check wrapped errors
errors.Is(err, ErrNotFound)      // check specific error
errors.As(err, &myErr)           // extract custom type
```

---

## Concurrency

```go
// Goroutine — lightweight thread
go someFunc()
go func() { /* inline */ }()

// Channel — safe data pipe
ch := make(chan string)        // unbuffered
ch := make(chan string, 5)     // buffered
ch <- "value"                  // send
val := <-ch                    // receive
close(ch)                      // close (for range)
for v := range ch { }          // drain channel

// WaitGroup — wait for goroutines
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); doWork() }()
wg.Wait()

// Mutex — protect shared state
var mu sync.Mutex
mu.Lock(); counter++; mu.Unlock()

// Select — wait on multiple channels
select {
case msg := <-ch1: fmt.Println(msg)
case msg := <-ch2: fmt.Println(msg)
case <-time.After(1 * time.Second): fmt.Println("timeout")
}
```

---

## JSON

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Pass  string `json:"-"`           // never in JSON
    Bio   string `json:"bio,omitempty"` // skip if empty
}

// Go → JSON
data, err := json.Marshal(user)
pretty, _ := json.MarshalIndent(user, "", "  ")

// JSON → Go
var u User
json.Unmarshal([]byte(jsonStr), &u)

// Stream encode/decode (for HTTP)
json.NewEncoder(w).Encode(v)           // write to ResponseWriter
json.NewDecoder(r.Body).Decode(&v)     // read from Request.Body
```

---

## HTTP Server

```go
func handler(w http.ResponseWriter, r *http.Request) {
    r.Method                          // "GET", "POST"...
    r.URL.Path                        // "/users"
    r.URL.Query().Get("id")           // ?id=42
    json.NewDecoder(r.Body).Decode(&v) // parse body

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)      // 200
    json.NewEncoder(w).Encode(data)
}

mux := http.NewServeMux()
mux.HandleFunc("/path", handler)
http.ListenAndServe(":8080", mux)

// Status codes
http.StatusOK           // 200
http.StatusCreated      // 201
http.StatusBadRequest   // 400
http.StatusUnauthorized // 401
http.StatusNotFound     // 404
http.StatusInternalServerError // 500
```

## Middleware

```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // before
        next.ServeHTTP(w, r)
        // after
    })
}
mux.Handle("/path", Logger(http.HandlerFunc(myHandler)))
```

---

## Environment Variables

```go
os.Getenv("KEY")             // value or ""
val, ok := os.LookupEnv("KEY") // value + exists bool
os.Setenv("KEY", "value")

// Best practice: load into a Config struct at startup
type Config struct { Port string; DBUrl string }
cfg := Config{
    Port:  getEnv("PORT", "8080"),
    DBUrl: getEnv("DATABASE_URL", ""),
}


```
