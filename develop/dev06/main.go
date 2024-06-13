package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Ожидание записи в канал
*/
func WaitReading() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("1. Выполнение горутины")
		for {
			select {
			case <-ch:
				fmt.Println("1. Остановка горутины")
				return
			}
		}
	}()
	ch <- 10
}

/*
С помощью завершение всех горутин из Wait Group
*/
func WaitEndWG() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	go func() {
		fmt.Println("2. Выполнение горутины")
		fmt.Println("2. Остановка горутины")
		wg.Done()
	}()
}

/*
С помощью тикера по истечению времени
*/
func WaitTimeTicker() {
	ticker := time.NewTicker(time.Second)
	ch := make(chan int)
	go func() {
		fmt.Println("3. Выполнение горутины")
		for {
			select {
			case <-ticker.C:
				ch <- 5
				fmt.Println("3. Остановка горутины")
				return
			}
		}
	}()
	<-ch
}

/*
	Остановка с помощью контекста withCancel
	Вызов cancel закрывает канал, возвращаемый методом контекста Done()
*/

func WaitContextEnding() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go func(ctx context.Context) {
		fmt.Println("4. Выполнение горутины")
		for {
			select {
			case <-ctx.Done():
				ch <- 1
				fmt.Println("4. Остановка горутины")
				return
			}
		}
	}(ctx)
	cancel()
	<-ch
}

func main() {
	WaitReading()
	WaitEndWG()
	WaitTimeTicker()
	WaitContextEnding()
}
