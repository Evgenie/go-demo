package file

import (
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

func (db *JsonDB) Write(content []byte) error {
	file, err := os.Create(db.fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		return err
	}
	return nil
}

func (db *JsonDB) Read() ([]byte, error) {
	return os.ReadFile(db.fileName)
}
