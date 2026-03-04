package main

import (
	"fmt"
	"strings"
)

// **Описание**: Создайте интерфейс Processor с методом Process([]byte) []byte и реализуйте тип UpperCaseProcessor, который преобразует все буквы в верхний регистр
//
// **Входные данные**: Встроенные тестовые данные: []byte("hello world"), []byte("go programming"), []byte("test data")
//
// **Выходные данные**: Результаты преобразования для каждого набора данных через интерфейс Processor
//
// **Ограничения**:
// - Используйте только стандартные возможности Go
// - Processor должен содержать только метод Process([]byte) []byte
// - UpperCaseProcessor должен преобразовывать все ASCII буквы в верхний регистр
// - Демонстрируйте работу через переменную типа Processor
// - Выводите исходные и преобразованные данные в консоль
//
// **Примеры**:
// Input: []byte("hello world")
// Output: []byte("HELLO WORLD")
//
// Входные данные: []byte("go programming")
// Output: []byte("GO PROGRAMMING")

// Ваш код здесь
type Processor interface {
	Process([]byte) []byte
}

type UpperCaseProcessor struct{}

func (UpperCaseProcessor) Process(data []byte) []byte {
	return []byte(strings.ToUpper(string(data)))
}

func main() {
	// Ваш код здесь
	var upperCaseProcessor Processor = &UpperCaseProcessor{}
	fmt.Printf("%s\n", upperCaseProcessor.Process([]byte("hello world")))
	fmt.Printf("%s\n", upperCaseProcessor.Process([]byte("go programming")))
	fmt.Printf("%s\n", upperCaseProcessor.Process([]byte("test data")))
}

