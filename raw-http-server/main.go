package main

import (
	"fmt"
	"log"
	"net"
)

// ============================================================
// MAIN — the raw TCP listener
//
// This is where it all begins. No net/http, no http.Server,
// no mux. Just a TCP socket opened on a port.
//
// The flow:
//   net.Listen("tcp", ":8081")   ← open a TCP port
//   listener.Accept()            ← block until a client connects
//   go handleConnection(conn)    ← hand it off to a goroutine
//
// That's it. Every browser tab, every curl command, every API
// call opens a new TCP connection. We spawn a goroutine for
// each one — Go's scheduler handles the rest.
// ============================================================

func main() {
	addr := ":8081"

	// Open a TCP socket on port 8081
	// net.Listen returns a net.Listener — an object that can
	// Accept() incoming connections.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to bind port:", err)
	}
	defer listener.Close()

	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║      TCP HTTP Server — No frameworks    ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║  Listening on http://localhost:8081          ║")
	fmt.Println("║                                              ║")
	fmt.Println("║  GET  /                                      ║")
	fmt.Println("║  GET  /hello                                 ║")
	fmt.Println("║  GET  /about                                 ║")
	fmt.Println("║  GET  /static/index.html                     ║")
	fmt.Println("║                                              ║")
	fmt.Println("║  Watch the terminal — raw request lines      ║")
	fmt.Println("║  will appear as bytes come off the wire!     ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	// ── Accept loop ────────────────────────────────────────
	// This loop runs forever. Each iteration:
	//   1. Blocks until a TCP client connects
	//   2. Gets a net.Conn — a bidirectional byte stream
	//   3. Hands it to a goroutine and immediately loops back
	//      to accept the next connection
	//
	// Goroutines are cheap (~4KB stack vs ~1MB for OS threads),
	// so spawning one per connection is perfectly fine in Go.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue // don't crash the whole server on one bad connection
		}

	
		go handleConnection(conn)
	}
}
