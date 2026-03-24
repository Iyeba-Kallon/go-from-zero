package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Request holds the parsed fields of an incoming HTTP request.
type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	RawLine string
}

// parseRequest reads the HTTP request off the TCP connection line by line.
// HTTP requests are plain text: a request line, then headers, then a blank line.
func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// First line is the request line, e.g. "GET /hello HTTP/1.1"
	rawLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading request line: %w", err)
	}
	rawLine = strings.TrimSpace(rawLine)

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

	// Read headers until the blank line that ends them.
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		colonIdx := strings.Index(line, ":")
		if colonIdx > 0 {
			key := strings.TrimSpace(line[:colonIdx])
			val := strings.TrimSpace(line[colonIdx+1:])
			req.Headers[key] = val
		}
	}

	return req, nil
}

// statusText maps status codes to their reason phrases.
var statusText = map[int]string{
	200: "OK",
	201: "Created",
	404: "Not Found",
	405: "Method Not Allowed",
	500: "Internal Server Error",
}

// Response holds the data needed to send an HTTP response.
type Response struct {
	StatusCode  int
	ContentType string
	Body        []byte
}

// write sends the response as raw HTTP/1.1 bytes over the TCP connection.
func (resp *Response) write(conn net.Conn) {
	reason, ok := statusText[resp.StatusCode]
	if !ok {
		reason = "Unknown"
	}

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
