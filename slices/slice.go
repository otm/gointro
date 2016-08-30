package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)
	fmt.Printf("len: %v, cap: %v\n", len(slice), cap(slice))
	slice = append(slice, 10, 20)
	fmt.Printf("len: %v, cap: %v\n", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Println(i, ":", v)
	}
}
