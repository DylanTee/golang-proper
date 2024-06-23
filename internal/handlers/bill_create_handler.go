package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type BillCreateRequest struct {
	UserId            string `json:"userId"`
	ProductId         string `json:"productId"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Amount            string `json:"amount"`
	ReferenceOne      string `json:"referenceOne"`
	ReferenceOneLabel string `json:"referenceOneLabel"`
}

func BillCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BillCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Println(req)

	apiURL := "https://www.billplz.com/api/v3/bills"

	data := url.Values{}
	data.Set("collection_id", "yec1sgag")
	data.Set("description", req.Description)
	data.Set("email", req.Email)
	data.Set("name", req.Name)
	data.Set("amount", req.Amount)
	data.Set("reference_1_label", req.ReferenceOneLabel)
	data.Set("reference_1", req.ReferenceOne)
	data.Set("callback_url", "https://pr0per.vercel.app/")

	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.SetBasicAuth("ab8ecbdc-51ec-4d5f-8658-c588a3930c06", "")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
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
	json.NewEncoder(w).Encode(body)
}
