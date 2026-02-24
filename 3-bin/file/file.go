package file

import (
	"errors"
	"os"
	"strings"
)

func ReadJson(name string) (*[]byte, error) {
	if !strings.Contains(name, ".json") {
		return nil, errors.New("Extension should be JSON")
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
