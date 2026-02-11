package main

import (
	"fmt"
	"strings"
)

// **Описание**: Создайте структуру Document с полями Title (строка), Content (строка) и WordCount (число), затем реализуйте функцию, которая принимает указатель на Document и подсчитывает количество слов в поле Content, обновляя поле WordCount
//
// **Входные данные**: Встроенные данные в коде (значения для инициализации структуры с заполненным Content и нулевым WordCount)
//
// **Выходные данные**: Вывод информации о документе до и после подсчета в формате "Title: [название], Content: [содержимое], WordCount: [количество]"
//
// **Ограничения**:
// - Используйте только базовые типы данных Go
// - Название должно быть непустой строкой
// - Содержимое может быть любой строкой (включая пустую)
// - Количество слов должно быть неотрицательным числом
// - Слова разделяются пробелами
// - Пустые строки и строки только из пробелов содержат 0 слов
//
// **Примеры**:
// Input: Title: "Report", Content: "Hello world from Go", WordCount: 0
// Output:
// Before: Title: Report, Content: Hello world from Go, WordCount: 0
// After: Title: Report, Content: Hello world from Go, WordCount: 4
//
// Входные данные: Title: "Note", Content: "Programming is fun", WordCount: 0
// Output:
// Before: Title: Note, Content: Programming is fun, WordCount: 0
// After: Title: Note, Content: Programming is fun, WordCount: 3

type Document struct {
	Title     string
	Content   string
	WordCount int
}

func (document *Document) countContentWords() {
	document.WordCount = len(strings.Fields(document.Content))
}

func newDocument(title string, content string, wordCount int) (*Document, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	if wordCount < 0 {
		return nil, fmt.Errorf("content should be non-negative")
	}

	return &Document{
		Title:     title,
		Content:   content,
		WordCount: wordCount,
	}, nil
}

func outputDocument(document *Document) string {
	return fmt.Sprintf("Title: %s, Content: %s, WordCount: %d", document.Title, document.Content, document.WordCount)
}

func outputResult(title string, content string, wordCount int) {
	document, err := newDocument(title, content, wordCount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Before:", outputDocument(document))
	document.countContentWords()
	fmt.Println("After:", outputDocument(document))
}

func main() {
	// Ваша реализация
	outputResult("Report", "Hello world from Go", 0)
	outputResult("Note", "Programming is fun", 0)
	outputResult("Note", "", 0)
	outputResult("Note", "     ", 0)
}
