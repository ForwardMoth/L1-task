package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Printf("a=%v b=%v\n", a, b)
	a, b = b, a
	fmt.Printf("a=%v b=%v\n", a, b)
}
