package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"block-cipher-webapp/backend/cipher"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

// POST /encrypt
func encryptHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)

	// Get request body
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("reqBody", reqBody)

	var message = reqBody["message"].(string)
	var key = reqBody["key"].(string)
	var mode = reqBody["mode"].(string)

	// Encrypt the plaintext
	result := cipher.Encrypt(message, key, mode)
	fmt.Println(result)
	
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// POST /decrypt
func decryptHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)

	// Get request body
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("reqBody", reqBody)

	var message = reqBody["message"].(string)
	var key = reqBody["key"].(string)
	var mode = reqBody["mode"].(string)

	// Decrypt the plaintext
	result := cipher.Decrypt(message, key, mode)
	fmt.Println(result)
	
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}