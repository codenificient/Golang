package main

import "fmt"

func main() {
	fmt.Println("We will see pointers now")

	x := 1
	p := &x

	fmt.Println("The value of the integer pointer is", p)
}