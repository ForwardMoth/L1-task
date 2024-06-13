package main

import "fmt"

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
*/

/*
Зададим строку локальной переменной
*/
func someFunc() []byte {
	hugeString := "Every type that is a member of the type set of an interface implements that interface. " +
		"Any given type may implement several distinct interfaces. For instance, all types implement the empty " +
		"interface which stands for the set of all (non-interface) types"

	justString := make([]byte, 100)
	copy(justString, hugeString)
	return justString
}

func main() {
	res := someFunc()
	fmt.Printf("%s", res)
}
