package account

import (
	"demo/account/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

const vaultFileName = "data.json"

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func InitVault() *Vault {
	db := files.NewJsonDB(vaultFileName)
	file, err := db.Read()
	var vault Vault
	if err != nil {
		vault = Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
		vault.WriteToJSON()
		return &vault
	}
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Ошибка чтения файла, создан новый")
		vault = Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
		vault.WriteToJSON()
		return &vault
	}
	return &vault
}

func (vault *Vault) WriteToJSON() {
	bytes, err := vault.ToBytes()

	if err != nil {
		color.Red(err.Error())
		return
	}

	db := files.NewJsonDB(vaultFileName)
	db.Write(bytes)
}

func (vault *Vault) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.UpdatedAt = time.Now()

	vault.WriteToJSON()
}

func (vault *Vault) DeleteAccount(url string) {
	for i, acc := range vault.Accounts {
		if acc.Url == url {
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			vault.UpdatedAt = time.Now()
			vault.WriteToJSON()
			color.HiGreen("Аккаунт успешно удален")
			return
		}
	}
	color.HiRed("Аккаунт не найден")
}

func (vault *Vault) OutputAccountList() {
	for i, acc := range vault.Accounts {
		color.HiYellow("%d. %+v\n", i+1, acc)
	}
}

func (vault *Vault) OutputAccount(login string) {
	for _, acc := range vault.Accounts {
		if acc.Login == login {
			color.HiGreen("Аккаунт:")
			acc.Output()
			return
		}
	}
	color.HiRed("Аккаунт не найден")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountsByURL(url string) *[]*Account {
	var targetAccounts []*Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			targetAccounts = append(targetAccounts, &acc)
		}
	}
	return &targetAccounts
}
