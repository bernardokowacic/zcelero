package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"zcelero/entity"
	"zcelero/repository"

	"github.com/google/uuid"
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

func (t *TextManagementService) Get(textId, privateKeyString, password string) (string, error) {
	pk := strings.NewReader(privateKeyString)
	pemBytes, err := ioutil.ReadAll(pk)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(pemBytes)
	bytePK, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		return "", err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(bytePK)
	if err != nil {
		return "", err
	}

	data, err := t.TextManagementRepository.Load(textId)
	if err != nil {
		return "", err
	}

	decriptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		return "", err
	}

	return string(decriptedData), nil
}

func (t *TextManagementService) Insert(text entity.TextManagement) (entity.TextManagement, error) {
	if *text.Encryption {
		var err error
		text.TextData, text.PrivateKey, err = encryptMessage(text.KeySize, text.TextData, text.PrivateKeyPassword)
		if err != nil {
			return entity.TextManagement{}, err
		}
	}

	text.Uuid = generateUuid()

	err := t.TextManagementRepository.Save(text.Uuid, text.TextData)
	if err != nil {
		return entity.TextManagement{}, err
	}

	return text, nil
}

func encryptMessage(keySize uint64, textData string, privateKeyPassword string) (string, string, error) {
	randReader := rand.Reader
	privatekey, err := rsa.GenerateKey(randReader, int(keySize))
	if err != nil {
		return "", "", err
	}

	publicKey := &privatekey.PublicKey
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		return "", "", err
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), randReader, publicKey, []byte(textData), nil)
	if err != nil {
		fmt.Printf("error when encrypting data: %s \n", err)
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
		return "", "", err
	}

	return string(ciphertext), string(pem.EncodeToMemory(block)), nil
}

func generateUuid() string {
	return uuid.New().String()
}
