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
	"io"
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

		privateKey, err := decryptPrivateKey(privateKeyString, password)
		if err != nil {
			log.Error().Msg(err.Error())
			return "", err
		}

		message, err = decryptMessage(privateKey, content)
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
		randReader := rand.Reader
		publicKey, privateKey, err := generatePairKey(randReader, text.KeySize, text.PrivateKeyPassword)
		if err != nil {
			log.Error().Msg(err.Error())
			return entity.TextManagement{}, err
		}

		encodedMessage, err = encryptMessage(randReader, publicKey, text.TextData)
		if err != nil {
			log.Error().Msg(err.Error())
			return entity.TextManagement{}, err
		}

		log.Debug().Msg("Encoding into base64")

		text.PrivateKey = privateKey
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

func generatePairKey(randReader io.Reader, keySize uint64, privateKeyPassword string) (*rsa.PublicKey, string, error) {
	privatekey, _ := rsa.GenerateKey(randReader, int(keySize))
	publicKey := &privatekey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	block, err := x509.EncryptPEMBlock(randReader, block.Type, block.Bytes, []byte(privateKeyPassword), x509.PEMCipherAES256)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "", err
	}

	return publicKey, string(pem.EncodeToMemory(block)), nil
}

func encryptMessage(randReader io.Reader, publicKey *rsa.PublicKey, textData string) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), randReader, publicKey, []byte(textData), nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return ciphertext, nil
}

func decryptPrivateKey(privateKeyString string, password string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyString))
	bytePK, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(bytePK)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return privateKey, nil
}

func decryptMessage(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	decriptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return string(decriptedData), nil
}
