package repository

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
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

// Save saves the file into folder
func (t *textManagementRepositoryStruct) Save(fileName string, content string) error {
	log.Debug().Msg("Creating file")

	file, err := os.Create(fmt.Sprintf("storage/%s.json", fileName))
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
	data, err := os.ReadFile(fmt.Sprintf("storage/%s.json", fileName))
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return data, nil
}
