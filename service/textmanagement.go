package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"
	"zcelero/entity"
	"zcelero/repository"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type TextManagementServiceInteface interface {
	Get(textId, privateKey, password string) (string, error)
	Insert(text entity.TextManagement) (entity.TextManagement, error)
}

type TextManagementService struct {
	TextManagementRepository repository.TextManagementInterface
}

func NewService(textManagementRepository repository.TextManagementInterface) TextManagementServiceInteface {
	return &TextManagementService{TextManagementRepository: textManagementRepository}
}

// Get load the file content and decrypt it if necessary
func (t *TextManagementService) Get(textId, privateKeyString, password string) (string, error) {
	log.Debug().Msg("Loading message from file")

	encrypted := false
	fileName := textId
	if !t.TextManagementRepository.CheckFileExists(textId) {
		log.Debug().Msg("Checking if file is encrypted")

		fileName = "encrypted_" + textId
		encrypted = true
	}

	log.Debug().Msg("Opening file")

	data, err := t.TextManagementRepository.Load(fileName)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	message := string(data)
	if encrypted {
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

		log.Debug().Msg("Decrypting message")

		message, err = decryptMessage(privateKeyString, password, data)
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

	text.Uuid = generateUuid()
	fileName := text.Uuid

	if *text.Encryption {
		log.Debug().Msg("Encrypting message")

		fileName = "encrypted_" + text.Uuid

		var err error
		text.TextData, text.PrivateKey, err = encryptMessage(text.KeySize, text.TextData, text.PrivateKeyPassword)
		if err != nil {
			log.Error().Msg(err.Error())
			return entity.TextManagement{}, err
		}

		log.Debug().Msg("Encryption finished")
	}

	log.Debug().Msg("Saving data into file")

	err := t.TextManagementRepository.Save(fileName, text.TextData)
	if err != nil {
		log.Error().Msg(err.Error())
		return entity.TextManagement{}, err
	}

	log.Debug().Msg("Message saved successfully")

	return text, nil
}

func encryptMessage(keySize uint64, textData string, privateKeyPassword string) (string, string, error) {
	randReader := rand.Reader
	privatekey, err := rsa.GenerateKey(randReader, int(keySize))
	if err != nil {
		log.Error().Msg(err.Error())
		return "", "", err
	}

	publicKey := &privatekey.PublicKey
	if err != nil {
		log.Error().Msg(err.Error())
		return "", "", err
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), randReader, publicKey, []byte(textData), nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", "", err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	// Convert it to pem
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	// Encrypt the pem
	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(privateKeyPassword), x509.PEMCipherAES256)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", "", err
	}

	return string(ciphertext), string(pem.EncodeToMemory(block)), nil
}

func generateUuid() string {
	return uuid.New().String()
}

func decryptMessage(privateKeyString string, password string, data []byte) (string, error) {
	pk := strings.NewReader(privateKeyString)
	pemBytes, err := ioutil.ReadAll(pk)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	block, _ := pem.Decode(pemBytes)
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
