package main

import "net/http"

// ============================================================
// ROUTER
//
// Go's built-in http.ServeMux matches paths but doesn't
// distinguish between HTTP methods (GET vs POST).
//
// This file builds a tiny layer on top that dispatches by
// BOTH method AND path — which is exactly what every Go
// router framework (chi, gorilla/mux) does internally.
//
// Pattern:
//   router.add("METHOD /path", handlerFunc)
//
// Under the hood we register a single mux handler per path,
// then inside that handler we switch on r.Method.
// ============================================================

// route pairs an HTTP method with a handler function.
type route struct {
	method  string
	handler http.HandlerFunc
}

// Router wraps http.ServeMux and maps "METHOD /path" → handler.
type Router struct {
	mux    *http.ServeMux
	routes map[string][]route // key = path (e.g. "/users")
}

// NewRouter creates and returns a new Router.
func NewRouter() *Router {
	return &Router{
		mux:    http.NewServeMux(),
		routes: make(map[string][]route),
	}
}

// add registers a handler for a specific method and path.
// Call this via the convenience methods: GET, POST, etc.
func (r *Router) add(method, path string, h http.HandlerFunc) {
	// If this path hasn't been registered yet, add it to the mux
	if _, exists := r.routes[path]; !exists {
		// Capture 'path' for the closure
		p := path
		r.mux.HandleFunc(p, func(w http.ResponseWriter, req *http.Request) {
			// Find the right handler for this method
			for _, rt := range r.routes[p] {
				if rt.method == req.Method {
					rt.handler(w, req)
					return
				}
			}
			// No handler matched → 405 Method Not Allowed
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "method not allowed",
			})
		})
	}

	// Register the method → handler mapping
	r.routes[path] = append(r.routes[path], route{method: method, handler: h})
}

// GET registers a handler for GET requests.
func (r *Router) GET(path string, h http.HandlerFunc) { r.add(http.MethodGet, path, h) }

// POST registers a handler for POST requests.
func (r *Router) POST(path string, h http.HandlerFunc) { r.add(http.MethodPost, path, h) }

// ServeHTTP makes Router satisfy http.Handler so it can be
// passed directly to http.Server or wrapped with middleware.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
