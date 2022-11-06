package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"zcelero/entity"
	"zcelero/helper"
	"zcelero/repository"

	"github.com/rs/zerolog/log"
)

type fileContent struct {
	Content   string
	Encrypted bool
}

type TextManagementServiceInteface interface {
	Get(textId, privateKey, password string) (string, error)
	Insert(text entity.TextManagement) (entity.TextManagement, error)
}

type TextManagementService struct {
	TextManagementRepository repository.TextManagementInterface
	Helper                   helper.HelperInterface
}

func NewService(textManagementRepository repository.TextManagementInterface, helper helper.HelperInterface) TextManagementServiceInteface {
	return &TextManagementService{
		TextManagementRepository: textManagementRepository,
		Helper:                   helper,
	}
}

// Get load the file content and decrypt it if necessary
func (t *TextManagementService) Get(textId, privateKeyString, password string) (string, error) {
	log.Debug().Msg("Loading message from file")

	log.Debug().Msg("Opening file")

	data, err := t.TextManagementRepository.Load(textId)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	fileData := fileContent{}
	json.Unmarshal(data, &fileData)

	message := string(fileData.Content)
	if fileData.Encrypted {
		if privateKeyString == "" {
			err := errors.New("private key is required to read this file")
			log.Info().Msg(err.Error())
			return "", err
		}
		if password == "" {
			err := errors.New("password is required to read this file")
			log.Info().Msg(err.Error())
			return "", err
		}

		log.Debug().Msg("Decoding base64")

		content, _ := base64.StdEncoding.DecodeString(fileData.Content)

		log.Debug().Msg("Decrypting message")

		message, err = decryptMessage(privateKeyString, password, content)
		if err != nil {
			log.Error().Msg(err.Error())
			return "", err
		}
	}

	log.Debug().Msg("Message loaded successfully")

	return message, nil
}

// Insert encrypt the message if necessary and save into a file
func (t *TextManagementService) Insert(text entity.TextManagement) (entity.TextManagement, error) {
	log.Debug().Msg("Creating new file with message")

	text.Uuid = t.Helper.GenerateUuid()

	if *text.Encryption {
		log.Debug().Msg("Encrypting message")

		var err error
		var encodedMessage []byte
		encodedMessage, text.PrivateKey, err = t.Helper.EncryptMessage(text.KeySize, text.TextData, text.PrivateKeyPassword)
		if err != nil {
			log.Error().Msg(err.Error())
			return entity.TextManagement{}, err
		}

		log.Debug().Msg("Encoding into base64")

		text.TextData = base64.StdEncoding.EncodeToString(encodedMessage)

		log.Debug().Msg("Encryption finished")
	}

	log.Debug().Msg("Saving data into file")

	fileData := fileContent{
		Content:   text.TextData,
		Encrypted: *text.Encryption,
	}
	b, _ := json.Marshal(fileData)

	err := t.TextManagementRepository.Save(text.Uuid, string(b))
	if err != nil {
		log.Error().Msg(err.Error())
		return entity.TextManagement{}, err
	}

	log.Debug().Msg("Message saved successfully")

	return text, nil
}

func decryptMessage(privateKeyString string, password string, data []byte) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyString))
	bytePK, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(bytePK)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	decriptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return string(decriptedData), nil
}
