package main

import (
	"block-cipher-webapp/backend/cipher"
	"encoding/json"
	"fmt"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

// POST /goblockc
func handler(w http.ResponseWriter, r *http.Request) {
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
	var encryptOrDecrypt = reqBody["encryptOrDecrypt"].(string)
	var mode = reqBody["mode"].(string)

	// Process the message (encrypt or decrypt)
	result := cipher.GoBlockC(message, key, encryptOrDecrypt, mode)
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
	http.HandleFunc("/goblockc", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
