package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// handleConnection is called in its own goroutine for each incoming TCP connection.
// It parses the request, picks a response, and writes it back to the socket.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := parseRequest(conn)
	if err != nil {
		textResponse(500, "Internal Server Error").write(conn)
		log.Printf("parse error: %v", err)
		return
	}

	log.Printf("⟵  %s", req.RawLine)

	var resp *Response

	switch {
	case req.Path == "/" && req.Method == "GET":
		resp = textResponse(200,
			"Welcome to my raw TCP HTTP server!\n"+
				"No net/http — just bytes on the wire.\n\n"+
				"Try: /hello  /about  /static/index.html",
		)

	case req.Path == "/hello" && req.Method == "GET":
		resp = jsonResponse(200, `{"message":"Hello from raw TCP!","built_with":"net.Listen + go handleConnection(conn)"}`)

	case req.Path == "/about" && req.Method == "GET":
		resp = jsonResponse(200, `{"server":"raw-http-server","language":"Go","transport":"TCP","http_library":"none"}`)

	case strings.HasPrefix(req.Path, "/static/"):
		resp = serveStaticFile(req.Path)

	default:
		resp = jsonResponse(404, fmt.Sprintf(`{"error":"route %q not found"}`, req.Path))
	}

	resp.write(conn)
	log.Printf("⟶  %d %s", resp.StatusCode, req.Path)
}

// serveStaticFile reads a file from the ./static directory and returns it as a response.
func serveStaticFile(urlPath string) *Response {
	filename := strings.TrimPrefix(urlPath, "/static/")
	filePath := filepath.Join("static", filename)

	// Guard against directory traversal e.g. /static/../../etc/passwd
	cleanPath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanPath, "static") {
		return textResponse(403, "Forbidden")
	}

	data, err := os.ReadFile(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return textResponse(404, fmt.Sprintf("file %q not found", filename))
		}
		return textResponse(500, "could not read file")
	}

	return &Response{
		StatusCode:  200,
		ContentType: detectContentType(filename),
		Body:        data,
	}
}

// detectContentType returns a Content-Type string based on the file extension.
func detectContentType(filename string) string {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".html", ".htm":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	default:
		return "text/plain; charset=utf-8"
	}
}
