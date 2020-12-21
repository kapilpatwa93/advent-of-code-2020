package util

import (
	"io/ioutil"
	"os"
	"path"
)

//go:generate go run github.com/kapilpatwa93/advent-of-code-2020/util/generators/frequency-map --package=util --filePrefix=frequency_map_ --templatePath=${PWD}/util/generators/frequency-map/template.txt

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func WriteFile(data []byte, fileName, filePath string, append bool) error {
	fileParams := os.O_RDWR | os.O_CREATE
	if append {
		fileParams |= os.O_APPEND
	} else {
		fileParams |= os.O_TRUNC
	}

	f, err := os.OpenFile(path.Join(filePath, fileName), fileParams, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
