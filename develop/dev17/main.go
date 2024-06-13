package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func fillArray(N int) []int {
	min := 1
	max := 20
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = rand.Intn(max-min) + min
	}
	return append(nums, 10)
}

func binarySearch(nums []int, x int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == x {
			return mid
		}

		if nums[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	N := 9
	nums := fillArray(N)
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println("pos %v", binarySearch(nums, 10))
	fmt.Println("pos %v", binarySearch(nums, -10))
}
