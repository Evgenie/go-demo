package main

import (
	"demo/account/account"
	"demo/account/files"
	"demo/account/output"
	"fmt"
	"strings"
)

var mapFunc = map[string]func(){
	"1": createAccount,
	"2": outputAccountList,
	"3": findAccounts,
	"4": deleteAccount,
}

var vault = account.InitVault(files.NewJsonDB("data.json"))

func promptData[T any](data []T) string {
	fmt.Println("")
	lastIdx := len(data) - 1
	for i, v := range data {
		if i == lastIdx {
			fmt.Printf("%v: ", v)
		} else {
			fmt.Println(v)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func getLogin() string {
	return promptData([]string{"Введите логин"})
}
func getURL() string {
	return promptData([]string{"Введите url"})
}

func createAccount() {
	myAccount, err := account.NewAccount(
		getLogin(),
		promptData([]string{"Введите пароль"}),
		getURL(),
	)

	if err != nil {
		output.PrintErrors(err)
		return
	}
	vault.AddAccount(myAccount)
}

func outputAccountList() {
	vault.OutputAccountList()
}
func findAccounts() {
	requiredUrl := getURL()
	accounts := vault.FindAccounts(func(a *account.Account) bool {
		return strings.Contains(a.Url, requiredUrl)
	})

	if len(*accounts) == 0 {
		output.PrintErrors("Аккаунтов не найдено")
		return
	}

	for _, acc := range *accounts {
		acc.Output()
	}
}
func deleteAccount() {
	vault.DeleteAccount(getURL())
}

func manageAccounts() {
Menu:
	for {
		action := promptData([]string{
			"1. Создать аккаунт",
			"2. Показать список аккаунтов",
			"3. Найти аккаунт",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите действие",
		})

		actionFunc := mapFunc[action]
		if actionFunc == nil {
			break Menu
		}
		actionFunc()
	}
}

func main() {
	fmt.Println("__ Менеджер акаунтов __")
	manageAccounts()
}
