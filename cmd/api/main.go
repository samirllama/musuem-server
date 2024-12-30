package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samirllama/musuem-server/internal/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	server_instance := server.New() // Create server instance

	// Create channel for shutdown signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server instance in a goroutine
	go func() {
		log.Printf("Starting server...")
		err := server_instance.Start("3333")
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-done // Wait for shutdown
	log.Printf("Server shutting down..")
}
