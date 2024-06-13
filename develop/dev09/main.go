package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	chNums := make(chan int)
	chSquares := make(chan int)

	go ReadArray(numbers, chNums)
	go WriteSquare(chNums, chSquares)

	// читаем значения из 2 канала
	for num := range chSquares {
		fmt.Println(num)
	}
}

// запись чисел в первый канал
func ReadArray(numbers []int, ch chan int) {
	defer close(ch)
	for _, num := range numbers {
		ch <- num
	}
}

// чтение числе из 1 канала и запись во 2 канал
func WriteSquare(chNums chan int, chSquares chan int) {
	defer close(chSquares)
	for num := range chNums {
		chSquares <- num * num
	}
}
