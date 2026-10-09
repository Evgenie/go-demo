package account

import (
	"demo/account/encrypter"
	"demo/account/output"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Encrypter = encrypter.Encrypter

type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write([]byte)
}
type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc Encrypter
}

func InitVault(db Db, enc Encrypter) *VaultWithDb {
	encryptedFile, err := db.Read()
	vault := &VaultWithDb{
		Vault: Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		},
		db:  db,
		enc: enc,
	}
	if err == nil {
		err = json.Unmarshal(vault.enc.Decrypt(encryptedFile), &vault.Vault)
		if err == nil {
			fmt.Println("Количество аккаунтов", len(vault.Accounts))
			return vault
		}
	}
	output.PrintErrors("Ошибка чтения файла, создан новый")
	vault.WriteToFile()
	return vault
}

func (vault *VaultWithDb) WriteToFile() {
	bytes, err := vault.ToBytes()

	if err != nil {
		output.PrintErrors(err.Error())
		return
	}

	vault.db.Write(vault.enc.Encrypt(bytes))
}

func (vault *VaultWithDb) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, *account)
	vault.UpdatedAt = time.Now()

	vault.WriteToFile()
}

func (vault *VaultWithDb) DeleteAccount(url string) {
	for i, acc := range vault.Accounts {
		if acc.Url == url {
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			vault.UpdatedAt = time.Now()
			vault.WriteToFile()
			color.HiGreen("Аккаунт успешно удален")
			return
		}
	}
	output.PrintErrors("Аккаунт не найден")
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
	output.PrintErrors("Аккаунт не найден")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccounts(value string, checker func(*Account, string) bool) *[]*Account {
	var targetAccounts []*Account
	for _, acc := range vault.Accounts {
		if checker(&acc, value) {
			targetAccounts = append(targetAccounts, &acc)
		}
	}
	return &targetAccounts
}
