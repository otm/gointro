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

var c = counter{}
var jobs = make(chan Message, 3)

func worker(jobs chan Message) {
	for {
		select {
		case msg := <-jobs:
			log.Println("Reseived message:", msg.Text, msg.ID)
			time.Sleep(1 * time.Second)
			log.Println("Message processed:", msg.ID)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := c.inc()
	msg := Message{Text: "Do work", ID: id}

	select {
	case jobs <- msg:
		log.Println("Message sent", id)
	case <-time.After(5 * time.Millisecond):
		log.Println("Failed to send message", id)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(msg)
}

// for i in {1..10}; do curl localhost:8080; done
func main() {
	// Start worker
	go worker(jobs)

	// Setup web server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
