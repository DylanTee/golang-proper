package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {

	// Define HTTP server.
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/bill/create", billCreateHandler)

	fmt.Println("Starting server at :8080")

	// PORT environment variable is provided by Cloud Run.
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "version 1")
	// Create an instance of the response struct
	response := Response{Message: "3"}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct into JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}

func billCreateHandler(w http.ResponseWriter, r *http.Request) {

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// var req Request
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&req)
	// if err != nil {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	// response := Response{Message: fmt.Sprintf("Hello, %s!", req.Name)}

	apiURL := "https://www.billplz-sandbox.com/api/v3/bills"

	data := url.Values{}
	data.Set("collection_id", "wvrcysgb")
	data.Set("description", "proper money subscription")
	data.Set("email", "aksoonz@gmail.com")
	data.Set("name", "Dylan Tee")
	data.Set("amount", "100")
	data.Set("reference_1_label", "REF 1 LABEL")
	data.Set("reference_1", "REF1")
	data.Set("callback_url", "https://pr0per.vercel.app/")

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("7fb4b5ef-877d-4293-85d8-c27b11d42e79", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("%+v\n", resp)
	fmt.Println(string(body))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Request.Form)

}
