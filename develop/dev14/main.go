package main

import "fmt"

func checkType(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	case chan int, chan string, chan bool:
		fmt.Println("Chan")
	default:
		fmt.Println(v)
	}
}

func main() {
	var a int
	var str string
	var flag bool
	var ch chan int

	checkType(a)
	checkType(str)
	checkType(flag)
	checkType(ch)
}
