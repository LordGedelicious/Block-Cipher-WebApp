package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	testjson := []string{"test1", "test2", "test3"}

	enableCors(&w)
	js, err := json.Marshal(testjson)
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