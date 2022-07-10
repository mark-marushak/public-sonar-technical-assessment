package parser

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrInvalidPath   = errors.New("invalid path to json file")
	ErrInvalidFormat = errors.New("invalid format of file")
)

func ParseJson(filepath string, output interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return ErrInvalidPath
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&output)
	if err != nil {
		return ErrInvalidFormat
	}

	return err
}
