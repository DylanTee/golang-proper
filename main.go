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
	http.HandleFunc("/postMethod", postHandler)

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
	response := Response{Message: "Proper version 1"}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct into JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	// billplz.Init("staging", "7fb4b5ef-877d-4293-85d8-c27b11d42e79")

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

	// data := models.Bill{
	// 	CollectionId:    "pb4cttjm5",
	// 	Email:           "aksoonz@gmail.com",
	// 	Mobile:          "60174449716",
	// 	Name:            "Dylan Tee",
	// 	Amount:          1,
	// 	CallbackUrl:     "https://billplz.com",
	// 	Description:     "Test Bill 123",
	// 	DueAt:           "2023-06-24",
	// 	Reference1Label: "Bank code",
	// 	Reference1:      "BP-1234",
	// }

	// resp, err := billplz.GetBill("12345")
	// if len(err.Error()) > 0 {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", resp)

	// response := Response{Message: fmt.Sprintf("Hello, %s!", req.Name)}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)

	// Define the URL and the data
	apiURL := "https://www.billplz.com/api/v3/bills"

	// Data to be sent in the POST request
	data := url.Values{}
	data.Set("collection_id", "pb4cttjm5")
	data.Set("description", "Maecenas eu placerat ante.")
	data.Set("email", "aksoonz@gmail.com")
	data.Set("name", "Sara")
	data.Set("amount", "1")
	data.Set("reference_1_label", "Bank Code")
	data.Set("reference_1", "BP-FKR01")
	data.Set("callback_url", "http://example.com/webhook/")

	// Create a new request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Basic Auth
	req.SetBasicAuth("7fb4b5ef-877d-4293-85d8-c27b11d42e79", "")

	// Create a client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response body
	fmt.Println(string(body))
}
