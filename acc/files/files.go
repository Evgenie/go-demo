package files

import (
	"fmt"
	"os"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		fileName: name,
	}
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	return os.ReadFile(db.fileName)
}
