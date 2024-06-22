package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// Define HTTP server.
	http.HandleFunc("/", helloRunHandler)

	fmt.Println("Starting server at :8080")

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coming from Proper Golang!!!!!!")
}
