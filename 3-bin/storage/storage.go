package storage

import (
	"demo/bins/bins"
	"demo/bins/file"
	"encoding/json"
	"os"
	"time"
)

const fileName = "storage.json"

type Storage struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (storage *Storage) SaveToFile() error {
	data, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, os.ModePerm)
}

func ReadFromFile() (*Storage, error) {
	file, err := file.ReadJson(fileName)
	storage := Storage{}
	if err != nil {
		return &storage, err
	}
	err = json.Unmarshal(*file, &storage)
	if err != nil {
		return &storage, err
	}
	return &storage, nil
}
