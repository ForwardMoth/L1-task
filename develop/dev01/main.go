package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human // Встраивание структуры Human в структуру Action
}

func (h *Human) Show() { // Show Human
	fmt.Printf("Human {name: %v, age: %v}\n", h.Name, h.Age)
}

func main() {
	action := Action{
		Human{Name: "Ivan", Age: 26},
	}
	action.Show() // Вызов метода структуры Human, встроенного в структуру Action
}
