package main

import (
	"fmt"
	"os"
)

/*
Проверяем валидный ли индекс
Затем берём слайс от 0 до i-1 элемента, после добавляем слайс i+1 до конца
... - означает расширение до списка аргументов: a1, a2, a3, и т.д.
*/
func deleteFromSlice(nums []int, i int) ([]int, error) {
	if i < 0 || i >= len(nums) {
		return nums, fmt.Errorf("error index")
	}

	return append(nums[:i], nums[i+1:]...), nil
}

func main() {
	nums := []int{2, 4, 6, 8, 10, 12}
	i := 4

	res, err := deleteFromSlice(nums, i)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res)
}
