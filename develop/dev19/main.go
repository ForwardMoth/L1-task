package main

import (
	"fmt"
)

/*
	Преобразуем строк в slice rune, затем меняем руны между собой с помощью указателя и возвращаем строку
*/

func reverseString(s string) string {
	v := []rune(s)
	N := len(v)
	for i := 0; i < N/2; i++ {
		v[i], v[N-i-1] = v[N-i-1], v[i]
	}
	return string(v)
}

func main() {
	var s string
	fmt.Scanf("%s", &s)

	res := reverseString(s)
	fmt.Println(res)
}
