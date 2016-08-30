package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Message holds the worker message
type Message struct {
	Text string `json:"text"`
	ID   int    `json:"id,omitempty"`
}

type counter struct {
	mu sync.Mutex
	i  int
}

func (c *counter) inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.i++
	return c.i
}

var tokens = make(chan struct{}, 5)
var c = counter{}

func handler(w http.ResponseWriter, r *http.Request) {
	id := c.inc()

	go func() {
		select {
		case <-tokens:
			defer func() { tokens <- struct{}{} }()
			log.Println("Processing message:", id)
			time.Sleep(90 * time.Millisecond)
			log.Println("Message processed:", id)
		default:
			log.Println("Unable to aquire token:", id)
		}
	}()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(Message{Text: "Job", ID: id})
}

// for i in {1..10}; do curl localhost:8080; done
func main() {
	// initialize token store
	for i := 0; i < 5; i++ {
		tokens <- struct{}{}
	}

	// Setup web server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
