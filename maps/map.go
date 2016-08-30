package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["foo"] = "bar"
	m["hello"] = "world"
	fmt.Println(m["hello"])

	if v, ok := m["foo"]; !ok {
		fmt.Println("key does not excist: foo")
	} else {
		fmt.Println("foo:", v)
	}
}
