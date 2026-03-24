package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// ============================================================
// ROUTER / CONNECTION HANDLER
//
// handleConnection is called in its own goroutine for every
// incoming TCP connection. This is the core of the server loop.
//
// It does three things:
//   1. Parses the raw HTTP request bytes into a Request struct
//   2. Decides which response to send (routing)
//   3. Writes the raw HTTP response bytes back to the socket
//
// Each connection is closed when this function returns —
// we're using HTTP/1.0-style "one request per connection"
// to keep things simple and clear.
// ============================================================

func handleConnection(conn net.Conn) {
	defer conn.Close() // always close the socket when done

	// ── 1. Parse the raw request ───────────────────────────
	req, err := parseRequest(conn)
	if err != nil {
		// If we can't even parse it, send a 500 and bail
		textResponse(500, "Internal Server Error").write(conn)
		log.Printf("parse error: %v", err)
		return
	}

	// ── 2. Log the raw request line (proof we're reading bytes!) ──
	log.Printf("⟵  %s", req.RawLine)

	// ── 3. Route by path ───────────────────────────────────
	var resp *Response

	switch {

	// GET / — home page
	case req.Path == "/" && req.Method == "GET":
		resp = textResponse(200,
			"Welcome to my raw TCP HTTP server!\n"+
				"No net/http — just bytes on the wire.\n\n"+
				"Try: /hello  /about  /static/index.html",
		)

	// GET /hello
	case req.Path == "/hello" && req.Method == "GET":
		resp = jsonResponse(200, `{"message":"Hello from raw TCP!","built_with":"net.Listen + go handleConnection(conn)"}`)

	// GET /about
	case req.Path == "/about" && req.Method == "GET":
		resp = jsonResponse(200, `{"server":"raw-http-server","language":"Go","transport":"TCP","http_library":"none — hand-rolled"}`)

	// GET /static/* — serve files from the ./static directory
	case strings.HasPrefix(req.Path, "/static/"):
		resp = serveStaticFile(req.Path)

	// Anything else → 404
	default:
		resp = jsonResponse(404, fmt.Sprintf(`{"error":"route %q not found"}`, req.Path))
	}

	// ── 4. Write raw HTTP response bytes to socket ─────────
	resp.write(conn)
	log.Printf("⟶  %d %s %s", resp.StatusCode, statusText[resp.StatusCode], req.Path)
}

// serveStaticFile reads a file from disk and returns its
// contents wrapped in a raw HTTP response.
// This is what a production server like nginx does, but we're
// doing it ourselves in ~20 lines.
func serveStaticFile(urlPath string) *Response {
	// Strip the /static/ prefix to get the filename
	// e.g. "/static/index.html" → "index.html"
	filename := strings.TrimPrefix(urlPath, "/static/")

	// Build the path relative to where the binary runs
	filePath := filepath.Join("static", filename)

	// Guard against directory traversal attacks: "../../../etc/passwd"
	// Clean resolves ".." components; if the result doesn't start
	// with "static" we refuse to serve it.
	cleanPath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanPath, "static") {
		return textResponse(403, "Forbidden")
	}

	// Read the entire file into memory
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return textResponse(404, fmt.Sprintf("file %q not found", filename))
		}
		return textResponse(500, "could not read file")
	}

	// Detect Content-Type by file extension
	contentType := detectContentType(filename)

	return &Response{
		StatusCode:  200,
		ContentType: contentType,
		Body:        data,
	}
}

// detectContentType returns a basic Content-Type string based on
// the file extension. A real server would use more MIME types.
func detectContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
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
