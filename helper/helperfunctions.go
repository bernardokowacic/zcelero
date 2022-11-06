package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type HelperInterface interface {
	EncryptMessage(keySize uint64, textData string, privateKeyPassword string) ([]byte, string, error)
	GenerateUuid() string
	CreateFile(filePath string) (*os.File, error)
	ReadFile(filePath string) ([]byte, error)
	WriteFile(file *os.File, content string) (n int, err error)
}

type helperStruct struct{}

func NewHelper() HelperInterface {
	return &helperStruct{}
}

// EncryptMessage encrypt a message and returns the encrypted message and the private key to decode it
func (h *helperStruct) EncryptMessage(keySize uint64, textData string, privateKeyPassword string) ([]byte, string, error) {
	randReader := rand.Reader

	privatekey, _ := rsa.GenerateKey(randReader, int(keySize))
	publicKey := &privatekey.PublicKey

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), randReader, publicKey, []byte(textData), nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "", err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(privateKeyPassword), x509.PEMCipherAES256)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, "", err
	}

	return ciphertext, string(pem.EncodeToMemory(block)), nil
}

// GenerateUuid generates a new UUID
func (h *helperStruct) GenerateUuid() string {
	return uuid.New().String()
}

func (h *helperStruct) CreateFile(filePath string) (*os.File, error) {
	return os.Create(filePath)
}

func (h *helperStruct) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (h *helperStruct) WriteFile(file *os.File, content string) (n int, err error) {
	return file.WriteString(content)
}
