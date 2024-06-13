package main

import "time"

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x)) // ожидание сигнала в канал по истечению времени
}

func main() {
	Sleep(5)
}
