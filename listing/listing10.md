<h3>Что выведет данная программа и почему?</h3>

```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```


Ответ:
```text
1
1
В update передается копия указателя и происходит изменение копии 
```
