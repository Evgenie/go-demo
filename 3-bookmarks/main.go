package main

import "fmt"

var bookmarks = map[string]string{}
var actions = []string{
	"Просмотр всех закладок",
	"Добавить закладку",
	"Удалить закладку",
	"Выход",
}

func hasBookmarks() bool {
	hasBookmarks := len(bookmarks) != 0
	if !hasBookmarks {
		fmt.Println("Список ваших закладок пуст.")
	}
	return hasBookmarks
}

func outputBookmark(name string, link string) {
	fmt.Printf("Название: %s, адрес: %s\n", name, link)
}

func outputMenu() {
	fmt.Println("")

	for i, action := range actions {
		fmt.Printf("%d. %s\n", i+1, action)
	}
}

func getAction() int {
	var action int
	outputMenu()
	fmt.Print("Выберите действие: ")
	fmt.Scan(&action)

	fmt.Println("")

	if action < 1 || action > 4 {
		fmt.Println("Ошибка. Повторите ввод.")

		action = getAction()
	}
	return action
}

func outputBookmarks() {
	if !hasBookmarks() {
		return
	}

	for name, link := range bookmarks {
		outputBookmark(name, link)
	}
}

func addBookmark() {
	var name string
	var link string
	var overwriteApproval string

	fmt.Print("Введите название: ")
	fmt.Scan(&name)

	if link, ok := bookmarks[name]; ok {
		fmt.Println("")
		fmt.Println("Закладка уже есть")
		outputBookmark(name, link)
		fmt.Print("Хотите перезаписать закладку? (y/n): ")
		fmt.Scan(&overwriteApproval)

		if overwriteApproval == "n" {
			return
		}
	}
	fmt.Print("Введите адрес: ")
	fmt.Scan(&link)
	bookmarks[name] = link

	fmt.Println("")

	outputBookmark(name, link)

	if overwriteApproval == "y" {
		fmt.Println("Закладка обновлена!")
	} else {
		fmt.Println("Закладка добавлена!")
	}
}

func deleteBookmark() {
	if !hasBookmarks() {
		return
	}

	var name string
	fmt.Print("Введите название для удаления закладки: ")
	fmt.Scan(&name)

	fmt.Println("")

	if link, ok := bookmarks[name]; ok {
		outputBookmark(name, link)
		fmt.Print("Уверенны, что хотите удалить закладку? (y/n): ")
		var approval string
		fmt.Scan(&approval)
		if approval == "y" {
			delete(bookmarks, name)
			fmt.Printf("Зкаладка %s удалена.\n", name)
		}
	} else {
		fmt.Printf("Закладки %s не существует, повторите ввод\n", name)
		deleteBookmark()
	}
}

func manageBookmarks() {
	action := getAction()

	switch action {
	case 1:
		outputBookmarks()
	case 2:
		addBookmark()
	case 3:
		deleteBookmark()
	default:
		return
	}
	manageBookmarks()
}

func main() {
	fmt.Println("___ Mенеджер закладок ___")
	manageBookmarks()
}
