package main

import (
	"fmt"
	"math/rand"
)

/*
	Запускаем воркеры.
	Воркер работает в бесконечном цикле, где происходит чтение из канала
	Запись в канал случайными числами происходит в основной горутине main
*/

func main() {
	fmt.Println("Введи  количество воркеров: ")
	var N int
	fmt.Scanf("%d", &N)

	ch := make(chan int)
	for i := 0; i < N; i++ {
		go worker(i, ch)
	}

	for {
		ch <- rand.Int()
	}
}

func worker(i int, ch chan int) {
	for {
		fmt.Printf("Read from %v worker data: %v\n", i, <-ch)
	}
}
