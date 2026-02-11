package main

import (
	"demo/bins/bins"
	"fmt"
	"time"
)

var NewBin = bins.NewBin

func main() {
	binList := bins.BinList{
		NewBin("1", true, time.Now(), "name1"),
		NewBin("12", true, time.Now(), "name2"),
		NewBin("3", true, time.Now(), "name3"),
		NewBin("4", true, time.Now(), "name4"),
	}
	for _, bin := range binList {
		fmt.Printf("%+v\n", *bin)
	}
}
