package main

import (
	"fmt"
	"strings"
)

// **Описание**: Создайте программу для поиска аккаунтов по имени пользователя в структуре Vault
//
// **Входные данные**: Готовая структура Vault с заполненным слайсом аккаунтов и строка для поиска
//
// **Выходные данные**: Вывод в консоль найденных аккаунтов в формате "Found: username at url" или "No accounts found" если ничего не найдено
//
// **Ограничения**:
// - Используйте пакет strings для сравнения имен пользователей
// - Поиск должен быть регистронезависимым
// - Выводите результат через fmt.Printf
// - Обработайте случай когда поисковая строка пустая
//
// **Примеры**:
// Input: Vault с аккаунтами ["alice", "bob", "Alice123"], поиск "alice"
// Output:
// Found: alice at example.com
// Found: Alice123 at test.org
//
// Входные данные: Vault с аккаунтами ["john", "mary"], поиск "peter"
// Output: No accounts found

type Account struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (account *Account) OutputAccount() {
	fmt.Printf("Found: %s at %s\n", account.Username, account.URL)
}

type Vault struct {
	Accounts []Account `json:"accounts"`
}

func (vault *Vault) FindAccounts(username string) {
	matched := false
	if username != "" {
		for _, account := range vault.Accounts {
			if strings.Contains(strings.ToLower(account.Username), strings.ToLower(username)) {
				account.OutputAccount()

				if !matched {
					matched = true
				}
			}
		}
	}
	if !matched {
		fmt.Println("No accounts found")
	}
}

func main() {
	// Готовые данные для тестирования
	vault := Vault{
		Accounts: []Account{
			{URL: "github.com", Username: "alice", Password: "pass1"},
			{URL: "gitlab.com", Username: "bob", Password: "pass2"},
			{URL: "example.com", Username: "Alice123", Password: "pass3"},
			{URL: "test.org", Username: "charlie", Password: "pass4"},
		},
	}

	searchQuery := "alice"

	// Ваш код здесь
	vault.FindAccounts(searchQuery)
}
