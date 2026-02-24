package storage

import (
	"demo/bins/bins"
	"demo/bins/file"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const fileName = "storage.json"

type Storage struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (storage *Storage) SaveToFile() {
	data, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFromFile() *Storage {
	file, err := file.ReadJson(fileName)
	storage := Storage{}
	if err != nil {
		fmt.Println(err)
		return &storage
	}
	err = json.Unmarshal(*file, &storage)
	if err != nil {
		fmt.Println(err)
		return &storage
	}
	return &storage
}
