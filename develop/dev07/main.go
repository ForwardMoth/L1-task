package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{4, 5, 121, 12, 1, 12}
	mp := make(map[int]int)
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(nums []int) {
		defer wg.Done()
		for i, n := range nums {
			mu.Lock()
			mp[i] = n
			mu.Unlock()
		}

	}(nums)
	wg.Wait()
	for k, v := range mp {
		fmt.Printf("Key %v Value %v\n", k, v)
	}
}
