package main

import "fmt"

func ChangeBits(n int64, i, bit int) int64 {
	if bit == 0 {
		return ChangeZeroBit(n, i)
	}

	return ChangeOneBit(n, i)
}

// сдвиг на i бит: 1 << 2 = 100 = 4
func BitShift(i int) int64 {
	return 1 << i
}

// Делаем побитовую инверсию устанавливаем все биты в 1, кроме i, помощью логического И сбрасываем бит
// 110 & ^(100) = 110 & 011 = 010 = 2
func ChangeZeroBit(n int64, i int) int64 {
	return n & ^BitShift(i)
}

// Логическое ИЛИ
// 110 | 100 = 110 = 6
func ChangeOneBit(n int64, i int) int64 {
	return n | BitShift(i)
}

func main() {
	var n int64 = 6 // 110
	i := 2

	newN := ChangeBits(n, i, 1)
	fmt.Println(newN)

	newN = ChangeBits(n, i, 0)
	fmt.Println(newN)
}
