package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("qwertyyuiopasdfghjklzxcvbnmQWERTYYUIOPASDFGHJKLZXCVBNM1234567890!@#$%^&*()_+{}[]<>./")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	rand.IntN(n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	(*acc).Password = string(res)
}
func (account *Account) Output() {
	color.Cyan(account.Login)
	color.Cyan(account.Password)
	color.Cyan(account.Url)
}

func NewAccount(login, password, urlInput string) (*Account, error) {
	_, err := url.ParseRequestURI(urlInput)
	switch {
	case login == "":
		return nil, errors.New("EMPTY_LOGIN")
	case err != nil:
		return nil, errors.New("INVALID_URL")

	default:
		nextAccount := Account{
			Login:     login,
			Password:  password,
			Url:       urlInput,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if password == "" {
			nextAccount.generatePassword(12)
		}
		return &nextAccount, nil
	}
}
