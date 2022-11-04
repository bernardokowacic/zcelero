package repository

import (
	"fmt"
	"os"
)

type TextManagementInterface interface {
	Save(fileName string, content string) error
	Load(fileName string) ([]byte, error)
}

type textManagementRepositoryStruct struct {
}

func NewRepository() TextManagementInterface {
	return &textManagementRepositoryStruct{}
}

func (t *textManagementRepositoryStruct) Save(fileName string, content string) error {
	file, err := os.Create(fmt.Sprintf("%s.txt", fileName))
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (t *textManagementRepositoryStruct) Load(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s.txt", fileName))
	if err != nil {
		return nil, err
	}

	return data, nil
}
