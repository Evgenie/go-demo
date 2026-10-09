package main

import (
	"demo/bins/bins"
	"demo/bins/file"
	"demo/bins/storage"
	"fmt"
	"time"
)

var NewBin = bins.NewBin

func main() {
	storage := storage.InitStorage(file.NewJsonDB("storage.json"))

	storage.AddBin(NewBin("1", true, time.Now(), "name1"))
	storage.AddBin(NewBin("12", true, time.Now(), "name2"))
	storage.AddBin(NewBin("3", true, time.Now(), "name3"))
	storage.AddBin(NewBin("4", true, time.Now(), "name4"))

	for _, bin := range storage.Bins {
		fmt.Printf("%+v\n", bin)
	}
}
