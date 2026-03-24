package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// ============================================================
// HTTP PARSING & WRITING — from raw bytes
//
// This is the layer that net/http does for you invisibly.
// We're doing it by hand so you can see exactly what HTTP
// looks like on the wire.
//
// An HTTP/1.1 request looks like this over the TCP socket:
//
//   GET /hello HTTP/1.1\r\n
//   Host: localhost:8081\r\n
//   User-Agent: curl/8.0\r\n
//   Accept: */*\r\n
//   \r\n
//
// An HTTP/1.1 response looks like this:
//
//   HTTP/1.1 200 OK\r\n
//   Content-Type: text/plain\r\n
//   Content-Length: 13\r\n
//   \r\n
//   Hello, world!
//
// The blank line (\r\n\r\n) separates headers from the body.
// That's the entire HTTP/1.1 format — it's just text!
// ============================================================

// Request holds the parsed fields of an incoming HTTP request.
type Request struct {
	Method  string            // "GET", "POST", etc.
	Path    string            // "/", "/hello", "/static/index.html"
	Version string            // "HTTP/1.1"
	Headers map[string]string // "Host" → "localhost:8081"
	RawLine string            // the original request line, great for logging
}

// parseRequest reads a raw HTTP request off the TCP connection
// and returns a parsed Request struct.
//
// It uses a bufio.Reader to read line-by-line from the socket.
// Each line ends with \r\n. The blank line signals end-of-headers.
func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// ── Step 1: Read the Request Line ────────────────────────
	// e.g. "GET /hello HTTP/1.1\r\n"
	rawLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading request line: %w", err)
	}
	rawLine = strings.TrimSpace(rawLine)

	// Split into three parts: method, path, version
	parts := strings.Fields(rawLine)
	if len(parts) < 3 {
		return nil, fmt.Errorf("malformed request line: %q", rawLine)
	}

	req := &Request{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: make(map[string]string),
		RawLine: rawLine,
	}

	// ── Step 2: Read Headers until blank line ─────────────────
	// Each line is "Key: Value\r\n". Stop at the blank line.
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // blank line = end of headers
		}

		// Split "Key: Value" on the first colon only
		colonIdx := strings.Index(line, ":")
		if colonIdx > 0 {
			key := strings.TrimSpace(line[:colonIdx])
			val := strings.TrimSpace(line[colonIdx+1:])
			req.Headers[key] = val
		}
	}

	return req, nil
}

// ============================================================
// RESPONSE WRITING
//
// We construct the raw HTTP response string and write it
// directly to the TCP socket as bytes.
// ============================================================

// statusText maps numeric status codes to their reason phrases.
var statusText = map[int]string{
	200: "OK",
	201: "Created",
	404: "Not Found",
	405: "Method Not Allowed",
	500: "Internal Server Error",
}

// Response holds everything needed to send an HTTP response.
type Response struct {
	StatusCode  int
	ContentType string
	Body        []byte
}

// write serialises the Response into raw HTTP/1.1 bytes and
// sends them down the TCP connection — nothing hidden, nothing magic.
func (resp *Response) write(conn net.Conn) {
	reason, ok := statusText[resp.StatusCode]
	if !ok {
		reason = "Unknown"
	}

	// Build the raw response manually, byte by byte
	header := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Type: %s\r\nContent-Length: %d\r\nConnection: close\r\n\r\n",
		resp.StatusCode,
		reason,
		resp.ContentType,
		len(resp.Body),
	)

	conn.Write([]byte(header))
	conn.Write(resp.Body)
}

// helpers so callers don't need to build Response themselves ──

func textResponse(status int, body string) *Response {
	return &Response{
		StatusCode:  status,
		ContentType: "text/plain; charset=utf-8",
		Body:        []byte(body),
	}
}

func htmlResponse(status int, body []byte) *Response {
	return &Response{
		StatusCode:  status,
		ContentType: "text/html; charset=utf-8",
		Body:        body,
	}
}

func jsonResponse(status int, body string) *Response {
	return &Response{
		StatusCode:  status,
		ContentType: "application/json",
		Body:        []byte(body),
	}
}
