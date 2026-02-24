package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

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

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
