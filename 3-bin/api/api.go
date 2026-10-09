package api

import (
	"demo/bins/config"
	"fmt"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", conf.Key)
}
