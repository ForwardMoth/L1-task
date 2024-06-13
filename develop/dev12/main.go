package main

import "fmt"

func newSet(nums []string) map[string]struct{} {
	s := map[string]struct{}{}
	for _, num := range nums {
		s[num] = struct{}{}
	}
	return s
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	wordsSet := newSet(words)

	fmt.Println(wordsSet)
}
