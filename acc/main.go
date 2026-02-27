package main

import (
	"demo/account/account"
	"demo/account/files"
	"fmt"

	"github.com/fatih/color"
)

var actions = map[string]int{
	"create": 1,
	"list":   2,
	"find":   3,
	"delete": 4,
	"exit":   5,
}

const actionError = "Не удалось распознать действие, повторите ввод."

var vault = account.InitVault(files.NewJsonDB("data.json"))

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
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

func createAccount() {
	myAccount, err := account.NewAccount(
		getLogin(),
		promptData("Введите пароль"),
		getURL(),
	)

	if err != nil {
		fmt.Println(err)
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
		color.Red("Аккаунтов не найдено")
		return
	}

	for _, acc := range *accounts {
		acc.Output()
	}
}
func deleteAccount() {
	vault.DeleteAccount(getURL())
}

func outputMenu() {
	fmt.Println("Выберите действие:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Показать список аккаунтов")
	fmt.Println("3. Найти аккаунт")
	fmt.Println("4. Удалить аккаунт")
	fmt.Println("5. Выход")
}

func getAction() int {
	var action int
	outputMenu()
	fmt.Scanln(&action)

	return action
}

func manageAccounts() {
Menu:
	for {
		switch getAction() {
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
