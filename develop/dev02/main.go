package main

import (
	"fmt"
	"sync"
)

/*
	Определяем группу горутин WaitGroup
	Функция FirstSolution блокируется на ожидание Wait, пока все горутины не будут выполнены в ней
	С помощью Add определяем количество горутин в группе для выполнения
	С помощью Done показываем завершение горутины
*/

func FirstSolution(nums []int) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(len(nums))
	for _, num := range nums {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(num)
	}
}

/*
	Создаем канал для передачи данных между горутинами
	По завершению работы с каналом закрываем его
	В цикле запускаем горутины, в которых происходит запись в канал.
	В это время канал на чтение блокируется, ожидая записи в канал
	Таким образом, происходит 5 раз запись в канал и 5 раз чтение из него
*/

func SecondSolution(nums []int) {
	ch := make(chan int)
	defer close(ch)
	for _, num := range nums {
		go func(i int) {
			ch <- i * i
		}(num)
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	FirstSolution(nums)
	fmt.Println()
	SecondSolution(nums)
}
