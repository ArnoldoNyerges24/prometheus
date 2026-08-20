package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondJSON buffers the response to ensure serialization errors occur before writing headers.
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func main() {
	fmt.Println("Hello, Bounty Hunter!")
}