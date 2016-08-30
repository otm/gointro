package main

import "fmt"

type celcius int

// String implements the Stringer interface
func (c celcius) String() string {
	return fmt.Sprintf("%v degrees warm", int(c))
}

func main() {
	temp := celcius(25)
	fmt.Println(temp)
	fmt.Printf("Type: %T", temp)
}
