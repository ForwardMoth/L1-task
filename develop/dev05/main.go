package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Создаём таймер с заданным количеством секунд
	Запускаем горутину. В ней работает бесконечный цикл.
	Для работы с каналом используем select. В случае, когда таймер достигает заданного количества секунд
	Канал закрывается и горутина прекращает работу
	Параллельно запущен цикл для чтения из горутины
*/

func main() {
	fmt.Println("Введи  количество секунд работы программы: ")
	var N int
	fmt.Scanf("%d", &N)

	ch := make(chan int)
	duration := time.Duration(N) * time.Second
	timer := time.NewTimer(duration)

	go func() {
		for {
			select {
			case <-timer.C:
				close(ch)
				return
			case ch <- rand.Int():
			}
		}
	}()

	for n := range ch {
		fmt.Printf("Чтение из канала значения %v\n", n)
	}

	fmt.Println("Конец работы программы")
}
