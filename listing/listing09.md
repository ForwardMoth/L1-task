<h3>Сколько существует способов задать переменную типа slice или map?</h3>

Ответ:
```go
mp := make(map[int]int)
list := make([]int)

mp := new(map[int]int)
list := new([]int)

mp := map[int]int{0: 1, 10: 8}
list := []int{1, 2, 3, 4}
```
