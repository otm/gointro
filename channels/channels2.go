package main

import (
	"fmt"
	"time"
)

func consumer(reqests, replies chan string) {
	for req := range reqests {
		fmt.Println(req)
		time.Sleep(100 * time.Millisecond) // simulating work
		replies <- "pong"
	}
	fmt.Println("consumer: done")
}

func producer() {
	requests := make(chan string)
	replies := make(chan string)

	go consumer(requests, replies)
	timer := time.After(1 * time.Second)

	for {
		select {
		case requests <- "ping":
		case msg := <-replies:
			fmt.Println(msg)
		case <-timer:
			fmt.Println("Done!!!")
			return
		}
	}
}

func main() {
	producer()
}
