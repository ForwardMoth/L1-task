package main

import "fmt"

// Создаем мапу, где ключ число, значение пустая структура и заполняем ими значения из массива (аналог set)
func newSet(nums []int) map[int]struct{} {
	s := map[int]struct{}{}
	for _, num := range nums {
		s[num] = struct{}{}
	}
	return s
}

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{3, 4, 5, 6, 7}

	s1 := newSet(a1)
	s2 := newSet(a2)

	if len(s2) > len(s1) {
		s1, s2 = s2, s1 // меняем множества, для итерации по множеству с меньшей длиной
	}

	s3 := newSet([]int{})
	for k, _ := range s2 {
		if _, ok := s1[k]; ok { // проверяем есть ли ключ из s2 в s1
			s3[k] = struct{}{}
		}
	}

	fmt.Println(s3)
}
