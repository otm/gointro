package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HelloWorld struct {
	Message string
	Date    time.Time `json:"date,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(HelloWorld{Message: "Hello World!", Date: time.Now()})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
