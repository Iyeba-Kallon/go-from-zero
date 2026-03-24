# go-from-zero 🚀

A collection of foundational systems projects I built during my first month learning Go. This repository tracks my journey from complete beginner to understanding what actually happens under the hood when a computer runs a process or fields a network request.

Read the full story on Medium: *[I Learned Go Last Month. I Built Three Things. Here's What Happened.](https://medium.com/@keepingupwithriya/i-learned-go-last-month-i-built-three-things-heres-what-happened-e951d9dae558)*

## Projects

### 1. Terminal Maze Game (`terminal-game/`)
A visual terminal escape maze game built entirely from scratch. You navigate a player (`@`) through walls (`#`) to reach a goal (`G`) using standard WASD controls.
- **Key Concepts:** Structs, game loops, pointer receivers (`*Game`), ANSI escape sequences, struct comparison.
- **Run it:**
  ```bash
  cd terminal-game
  go run main.go
  ```

### 2. Go Shell (`shell/`)
A working command-line interface capable of executing system commands, tracking the current working directory, and running internal commands like `cd`.
- **Key Concepts:** Infinite loops (`for { }`), buffer readers (`bufio`), process execution (`os/exec`), handling parent-child process relationships (like directory changes).
- **Run it:**
  ```bash
  cd shell
  go run main.go
  ```

### 3. Raw TCP HTTP Server (`raw-http-server/`)
A fully functioning HTTP server built without Go's `net/http` framework. It utilizes raw TCP sockets (`net.Listen`) and lightweight threads (`goroutines`) to manually parse incoming byte streams into HTTP requests and format standard HTTP/1.1 responses over the wire.
- **Key Concepts:** Sockets, `net.Conn`, TCP stream parsing, raw HTTP protocol formatting, MIME type detection, concurrency (`go handleConnection`).
- **Run it:**
  ```bash
  cd raw-http-server
  go run .
  # Then open http://localhost:8081/static/index.html to see the server serve its own UI!
  ```

---

## My Cheatsheet

Throughout this process, I compiled all of the Go mechanics I learned into a single, comprehensive reference guide. If you are learning Go, I highly recommend building one of these as you go.

 **View it here:** [notes/go-cheatsheet.md](notes/go-cheatsheet.md)

## Tech Stack
- **Language:** Go 1.22
- **External Dependencies:** Zero. The entire repository is built relying purely on Go's standard library.

## About the Author
I'm a final-year engineering student with an interest in systems programming. If you found this repository through my article, thank you for reading along!
