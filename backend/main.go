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
	var isEncrypt = reqBody["isEncrypt"].(bool)
	var mode = reqBody["mode"].(string)
	var isFile = reqBody["isFile"].(bool)

	// Process the message (encrypt or decrypt)
	result, timeElapsed := cipher.GoBlockC(message, key, mode, isEncrypt, isFile)
	fmt.Println(result, timeElapsed.String())

	response := map[string]interface{}{
		"result":      result,
		"timeElapsed": timeElapsed.String(),
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	return
}

func main() {
	http.HandleFunc("/goblockc", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
