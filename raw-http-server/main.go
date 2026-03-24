package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := ":8081"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to bind port:", err)
	}
	defer listener.Close()

	fmt.Println("raw-http-server listening on http://localhost" + addr)
	fmt.Println("routes: /  /hello  /about  /static/index.html")

	// Accept connections in a loop and handle each one in its own goroutine.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
