package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/10.0) * 10
		res[key] = append(res[key], temp)
	}
	fmt.Println(res)
}
