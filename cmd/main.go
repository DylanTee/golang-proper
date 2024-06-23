package main

import (
	"fmt"
	"golang-proper/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	// Define HTTP server.
	http.HandleFunc("/", handlers.VersionHandler)
	http.HandleFunc("/bill/create", handlers.BillCreateHandler)
	fmt.Println("Starting server at :8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
