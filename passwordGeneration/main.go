package main

import (
	"demo/account/files"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}
type TestAccount struct {
	account
}

func (acc *account) generatePassword(n int) {
	letterRunes := []rune("qwertyyuiopasdfghjklzxcvbnmQWERTYYUIOPASDFGHJKLZXCVBNM1234567890!@#$%^&*()_+{}[]<>./")
	res := make([]rune, n)
	rand.IntN(n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	(*acc).password = string(res)
}

func newAccount(login, password, urlInput string) (*account, error) {
	_, err := url.ParseRequestURI(urlInput)
	switch {
	case login == "":
		return nil, errors.New("EMPTY_LOGIN")
	case err != nil:
		return nil, errors.New("INVALID_URL")

	default:
		nextAccount := account{
			login:    login,
			password: password,
			url:      urlInput,
		}
		if password == "" {
			nextAccount.generatePassword(12)
		}
		return &nextAccount, nil
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func main() {
	myAccount, err := newAccount(promptData("Введите логин"), promptData("Введите пароль"), promptData("Введите url"))

	if err != nil {
		fmt.Println(err)
		return
	}

	files.ReadFile()

	fmt.Printf("%+v\n", myAccount)
}
