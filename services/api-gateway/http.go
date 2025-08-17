package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}
	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Print(err)
		return
	}
	log.Println(resp)
	fmt.Println("here...")
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to parse response", http.StatusBadRequest)
		return
	}
	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusCreated, response)
}
