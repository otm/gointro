package main

import "fmt"

func main() {
	requests := make(chan string)
	replies := make(chan string)

	go func() {
		fmt.Println(<-requests)
		replies <- "pong"
	}()

	requests <- "ping"
	msg := <-replies
	fmt.Println(msg)
}
