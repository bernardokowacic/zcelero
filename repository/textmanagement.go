package repository

import (
	"fmt"
	"zcelero/helper"

	"github.com/rs/zerolog/log"
)

type TextManagementInterface interface {
	Save(fileName string, content string) error
	Load(fileName string) ([]byte, error)
}

type textManagementRepositoryStruct struct {
	Helper helper.HelperInterface
}

func NewRepository(helper helper.HelperInterface) TextManagementInterface {
	return &textManagementRepositoryStruct{Helper: helper}
}

var fileLocation = "storage"

// Save saves the file into folder
func (t *textManagementRepositoryStruct) Save(fileName string, content string) error {
	log.Debug().Msg("Creating file")

	file, err := t.Helper.CreateFile(fmt.Sprintf("%s/%s.json", fileLocation, fileName))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	log.Debug().Msg("Writing data inside file")

	_, err = t.Helper.WriteFile(file, content)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

// Load reads the file into memory
func (t *textManagementRepositoryStruct) Load(fileName string) ([]byte, error) {
	log.Debug().Msg("Reading file")
	data, err := t.Helper.ReadFile(fmt.Sprintf("%s/%s.json", fileLocation, fileName))
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return data, nil
}
