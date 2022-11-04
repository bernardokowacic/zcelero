package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
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
	log.Debug().Msg("Creating file")

	file, err := os.Create(fmt.Sprintf("storage/%s.txt", fileName))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	log.Debug().Msg("Writing data inside file")

	_, err = file.WriteString(content)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

// Load reads the file into memory
func (t *textManagementRepositoryStruct) Load(fileName string) ([]byte, error) {
	log.Debug().Msg("Reading file")
	data, err := os.ReadFile(fmt.Sprintf("storage/%s.txt", fileName))
	if err != nil {
		log.Error().Msg(err.Error())
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
