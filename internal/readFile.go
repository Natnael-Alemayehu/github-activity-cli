package internal

import "os"

func ReadFile(fileName string) (*os.File, error) {
	data, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return data, nil
}
