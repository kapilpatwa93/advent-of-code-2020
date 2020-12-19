package util

import "io/ioutil"

func ReadInput(path string) ([]byte,error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data,nil
}
