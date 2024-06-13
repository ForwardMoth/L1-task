package main

import (
	"fmt"
	"sync"
)

/*
	Определяем группу горутин WaitGroup
	Функция FirstSolution блокируется на ожидание Wait, пока все горутины не будут выполнены в ней. После выводим результат
	С помощью Add определяем количество горутин в группе для выполнения
	С помощью Done показываем завершение горутины

	Для того чтобы получить правильный результат при изменении переменной, используется Mutex
	Т.е. для синхронизации между горутинами.
	С помощью блокировки и разблокировки гарантируется доступ только 1 горутины к переменной
	Тем самым получается ожидаемый результат
*/

func FirstSolution(nums []int) {
	wg := &sync.WaitGroup{}
	wg.Add(len(nums))
	var res = 0
	mu := &sync.Mutex{}
	for _, num := range nums {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			res += i * i
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(res)
}

/*
	Создаем канал для передачи данных между горутинами
	По завершению работы с каналом закрываем его
	В цикле запускаем горутины, в которых происходит запись в канал.
	В это время канал на чтение блокируется, ожидая записи в канал
	Таким образом, происходит 5 раз запись в канал и 5 раз чтение из него
	В переменную res прибавляем результат
*/

func SecondSolution(nums []int) {
	ch := make(chan int)
	defer close(ch)
	for _, num := range nums {
		go func(i int) {
			ch <- i * i
		}(num)
	}
	var res int
	for i := 0; i < len(nums); i++ {
		res += <-ch
	}
	fmt.Println(res)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	FirstSolution(nums)
	fmt.Println()
	SecondSolution(nums)
}
