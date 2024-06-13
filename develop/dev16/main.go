package main

import (
	"fmt"
	"math/rand"
)

func fillArray(N int) []int {
	min := 1
	max := 20
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = rand.Intn(max-min) + min
	}
	return nums
}

func partition(nums []int, l, r int) ([]int, int) {
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[r] = nums[r], nums[i]
	return nums, i
}

func quickSort(nums []int, l, r int) []int {
	if l < r {
		nums, p := partition(nums, l, r)
		quickSort(nums, l, p-1)
		quickSort(nums, p+1, r)
	}
	return nums
}

func main() {
	N := 6
	nums := fillArray(N)
	fmt.Println(nums)
	nums = quickSort(nums, 0, N-1)
	fmt.Println(nums)
}
