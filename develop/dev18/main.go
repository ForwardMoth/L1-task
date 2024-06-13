package main

import (
	"fmt"
	"sync"
)

// структура счетчика с mutex
type Counter struct {
	cnt int
	sync.Mutex
}

func main() {
	N := 1000
	ch := make(chan int) // использую канал, чтобы блокировать горутину main
	c := Counter{cnt: 0}

	go incriment(ch, N, &c)

	fmt.Println(<-ch)
}

// во время инкрементирования блокирую счетчик
func incriment(ch chan int, N int, c *Counter) {
	for i := 0; i < N; i++ {
		c.Lock()
		c.cnt++
		c.Unlock()
	}
	ch <- c.cnt // завершение горутины и разблокировка горутины main
}
