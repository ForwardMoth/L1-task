package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	N := len(words)
	for i := 0; i < N/2; i++ {
		words[i], words[N-i-1] = words[N-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	var ex1 = "snow dog sun"
	fmt.Println(reverseWords(ex1))
	var ex2 = "snow dog sun cat"
	fmt.Println(reverseWords(ex2))
}
