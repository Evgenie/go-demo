package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList = []*Bin

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}

func main() {
	binList := BinList{
		NewBin("1", true, time.Now(), "name1"),
		NewBin("12", true, time.Now(), "name2"),
		NewBin("3", true, time.Now(), "name3"),
		NewBin("4", true, time.Now(), "name4"),
	}
	for _, bin := range binList {
		fmt.Printf("%+v\n",*bin)
	}
}
