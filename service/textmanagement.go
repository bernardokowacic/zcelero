package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"zcelero/database/entity"
	"zcelero/repository"

	"github.com/google/uuid"
)

type TextManagementServiceInteface interface {
	Get()
	Insert(text entity.TextManagement) (entity.TextManagement, error)
}

type TextManagementService struct {
	TextManagementRepository repository.TextManagementInterface
}

func NewService(textManagementRepository repository.TextManagementInterface) TextManagementServiceInteface {
	return &TextManagementService{TextManagementRepository: textManagementRepository}
}

func (t *TextManagementService) Get() {

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

	fileName := fmt.Sprintf("%s.txt", text.Uuid)
	file, err := os.Create(fileName)
	if err != nil {
		return entity.TextManagement{}, err
	}

	_, err = file.WriteString(text.TextData)
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

	// Convert it to pem
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: ciphertext,
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
