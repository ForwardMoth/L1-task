<h3>Что выведет данная программа и почему?</h3>

```go
func main() {
    slice := []string{"a", "a"}

    func(slice []string) {
        slice = append(slice, "a")
        slice[0] = "b"
        slice[1] = "b"
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice)
}

```


Ответ:
```text
b b a 
a a 
В анонимной функции при использовании append создается новый слайс с новой ссылкой, в котором меняются значения
Поэтому старый слайс в main не изменяется
```