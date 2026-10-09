package storage

import (
	"demo/bins/bins"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Storage struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type StorageWithDb struct {
	Storage
	db Db
}

func InitStorage(db Db) *StorageWithDb {
	file, err := db.Read()
	storage := StorageWithDb{
		Storage: Storage{
			Bins:      []bins.Bin{},
			UpdatedAt: time.Now(),
		},
		db: db,
	}
	if err != nil {
		fmt.Println("Ошибка чтения файла, создан новый")
		storage.WriteToJSON()
		return &storage
	}
	err = json.Unmarshal(file, &storage.Storage)
	if err != nil {
		fmt.Println("Ошибка чтения файла, создан новый")
		storage.WriteToJSON()
		return &storage
	}
	return &storage
}

func (storage *StorageWithDb) WriteToJSON() {
	bytes, err := storage.ToBytes()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	storage.db.Write(bytes)
}

func (storage *StorageWithDb) AddBin(bin *bins.Bin) {
	storage.Bins = append(storage.Bins, *bin)
	storage.UpdatedAt = time.Now()

	storage.WriteToJSON()
}

func (storage *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(storage)

	if err != nil {
		return nil, err
	}
	return file, nil
}
