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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := cipher.Hello()
	enableCors(&w)

	js, err := json.Marshal(hello)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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

	// Encrypt the plaintext
	result := cipher.Encrypt(reqBody["plaintext"].(string))
	fmt.Println(result)
	// testjson := []string{"test1", "test2", "test3"}

	
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/encrypt", encryptHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}