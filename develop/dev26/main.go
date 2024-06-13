package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	s = strings.ToLower(s)
	var set = map[rune]struct{}{}

	for _, ch := range s {
		if _, ok := set[ch]; ok {
			return false
		}
		set[ch] = struct{}{}
	}
	return true
}

func main() {
	s1 := "asdfA"
	s2 := "abcdE"
	s3 := "abCdefAaf"
	s4 := "aabcd"

	fmt.Println(isUnique(s1))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
	fmt.Println(isUnique(s4))
}
