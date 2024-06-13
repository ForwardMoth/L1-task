<h3>Что выведет данная программа и почему?</h3>

```go
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(wg sync.WaitGroup, i int) {
            fmt.Println(i)
            wg.Done()
        }(wg, i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```


Ответ:
```text
Deadlock. Передача wg в качестве копии, а не указателя 
```
