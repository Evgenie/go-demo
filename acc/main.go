package main

import (
	"demo/account/account"
	"demo/account/files"
	"demo/account/output"
	"fmt"
	"strings"
)

var mapFunc = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": outputAccountList,
	"3": getFindAccounts(getURL, func(a *account.Account, url string) bool {
		return strings.Contains(a.Url, url)
	}),
	"4": getFindAccounts(getLogin, func(a *account.Account, login string) bool {
		return strings.Contains(a.Login, login)
	}),
	"5": deleteAccount,
}

func promptData(data ...string) string {
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
	return promptData("Введите логин")
}
func getURL() string {
	return promptData("Введите url")
}

func createAccount(vault *account.VaultWithDb) {
	myAccount, err := account.NewAccount(
		getLogin(),
		promptData("Введите пароль"),
		getURL(),
	)

	if err != nil {
		output.PrintErrors(err)
		return
	}
	vault.AddAccount(myAccount)
}

func outputAccountList(vault *account.VaultWithDb) {
	vault.OutputAccountList()
}
func getFindAccounts(getValue func() string, checker func(a *account.Account, s string) bool) func(vault *account.VaultWithDb) {
	return func(vault *account.VaultWithDb) {
		accounts := vault.FindAccounts(getValue(), checker)

		if len(*accounts) == 0 {
			output.PrintErrors("Аккаунтов не найдено")
			return
		}

		for _, acc := range *accounts {
			acc.Output()
		}
	}
}
func deleteAccount(vault *account.VaultWithDb) {
	vault.DeleteAccount(getURL())
}

func manageAccounts() {
	vault := account.InitVault(files.NewJsonDB("data.json"))
Menu:
	for {
		action := promptData(
			"1. Создать аккаунт",
			"2. Показать список аккаунтов",
			"3. Найти аккаунт по URL",
			"4. Найти аккаунт по логину",
			"5. Удалить аккаунт",
			"6. Выход",
			"Выберите действие",
		)

		actionFunc := mapFunc[action]
		if actionFunc == nil {
			break Menu
		}
		actionFunc(vault)
	}
}

func main() {
	fmt.Println("__ Менеджер акаунтов __")
	manageAccounts()
}
