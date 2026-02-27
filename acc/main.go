package main

import (
	"demo/account/account"
	"demo/account/files"
	"demo/account/output"
	"fmt"
)

var actions = map[string]string{
	"create": "1",
	"list":   "2",
	"find":   "3",
	"delete": "4",
	"exit":   "5",
}

const actionError = "Не удалось распознать действие, повторите ввод."

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
	accounts := vault.FindAccountsByURL(getURL())

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
		switch promptData([]string{
			"1. Создать аккаунт",
			"2. Показать список аккаунтов",
			"3. Найти аккаунт",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите действие",
		}) {
		case actions["create"]:
			createAccount()
			continue
		case actions["list"]:
			outputAccountList()
			continue
		case actions["find"]:
			findAccounts()
			continue
		case actions["delete"]:
			deleteAccount()
			continue
		case actions["exit"]:
			break Menu

		default:
			fmt.Println(actionError)
			continue
		}
	}
}

func main() {
	fmt.Println("__ Менеджер акаунтов __")
	manageAccounts()
}
