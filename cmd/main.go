package main

import (
	"log"
	"net/http"

	"github.com/puxin71/assignment"
)

const (
	// Our app's URI
	Addr = "localhost:3000"
)

func main() {
	db := assignment.NewDB()
	handler := assignment.NewHandler(db)

	// Configure server routes
	router := assignment.NewRouter(handler)
	http.Handle("/", router)
	server := http.Server{
		Handler: router,
		Addr:    Addr,
	}

	log.Println("Starting server and listens on " + Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe, error: %v", err)
	}
}
