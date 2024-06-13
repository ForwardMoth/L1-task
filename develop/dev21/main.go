package main

import "fmt"

// Интерфейс printer, который имплементирует метод getDocumentData
type Printer interface {
	getDocumentData()
}

// Адаптируемая структура
type WordDocument struct {
	Name string
	Len  int
}

// Метод адаптируемой структуры
func (d WordDocument) getWordDocumentData() {
	fmt.Printf("name document %v and len %v \n", d.Name, d.Len)
}

// Адаптер
type DocumentAdapter struct {
	document *WordDocument
}

// Реализация общего метода адаптером
func (adapter DocumentAdapter) getDocumentData() {
	adapter.document.getWordDocumentData()
}

// Создание адаптера
func newAdapter(doc *WordDocument) Printer {
	return &DocumentAdapter{document: doc}
}

func main() {
	document := WordDocument{Name: "test", Len: 4}
	adapter := newAdapter(&document)
	adapter.getDocumentData()
}
