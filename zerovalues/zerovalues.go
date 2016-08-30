package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var m map[string]string
	var slice []int
	fmt.Printf("%v %v %v %q, %#v, %#v\n", i, f, b, s, m, slice)
}
