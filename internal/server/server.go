package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/samirllama/musuem-server/internal/handlers"
)

// Server is a struct that holds our server's dependencies
// In this case, it only has a mux (request router) but we can add more things later
type Server struct {
	server *http.Server   // pointers let us modify the original value, not a copy
	mux    *http.ServeMux // ServeMux is Go's HTTP request router
}

// New is a constructor function that creates a new Server instance. The name "New" is a Go convention for constructor functions
// "*Server" means this function returns a pointer to a Server
// (pointers let us modify the original value, not a copy)
func New() *Server {
	mux := http.NewServeMux() // Create a new ServeMux (HTTP request router)

	// "&Server" creates a new Server and returns its memory address
	s := &Server{ // & operator gets memory address of something

		server: &http.Server{
			// Server will handle timeouts properly
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},

		mux: mux,
	}

	s.routes() // Set up all our routes

	s.server.Handler = s.mux // Set the router to our server

	return s // Return the server instance
}

// routes sets up all our HTTP endpoints
func (s *Server) routes() {
	log.Println("Setting up routes...")             // Log when routes are being registered
	s.mux.HandleFunc("/hell", handlers.HellHandler) // Just cus...

	s.mux.HandleFunc("/hello", handlers.HelloHandler)        // Basic hello endpoint
	s.mux.HandleFunc("/health", handlers.HealthCheckHandler) // Health check endpoint (commonly used in production)
}

func (s *Server) Start(port string) error {
	s.server.Addr = fmt.Sprintf(":%s", port) // Set the server address

	log.Printf("Server starting on port %s", port)

	// Start server
	return s.server.ListenAndServe()

}
