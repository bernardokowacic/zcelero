package repository

import (
	"errors"
	"fmt"
	"os"
)

type TextManagementInterface interface {
	Save(fileName string, content string) error
	Load(fileName string) ([]byte, error)
	CheckFileExists(fileName string) bool
}

type textManagementRepositoryStruct struct {
}

func NewRepository() TextManagementInterface {
	return &textManagementRepositoryStruct{}
}

// Save saves the file into folder
func (t *textManagementRepositoryStruct) Save(fileName string, content string) error {
	file, err := os.Create(fmt.Sprintf("storage/%s.txt", fileName))
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// Load reads the file into memory
func (t *textManagementRepositoryStruct) Load(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fmt.Sprintf("storage/%s.txt", fileName))
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CheckFileExists check if the file exists
func (t *textManagementRepositoryStruct) CheckFileExists(fileName string) bool {
	if _, err := os.Stat(fmt.Sprintf("storage/%s.txt", fileName)); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
