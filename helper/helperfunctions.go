package helper

import (
	"os"

	"github.com/google/uuid"
)

type HelperInterface interface {
	GenerateUuid() string
	CreateFile(filePath string) (*os.File, error)
	ReadFile(filePath string) ([]byte, error)
	WriteFile(file *os.File, content string) (n int, err error)
}

type helperStruct struct{}

func NewHelper() HelperInterface {
	return &helperStruct{}
}

// GenerateUuid generates a new UUID
func (h *helperStruct) GenerateUuid() string {
	return uuid.New().String()
}

// CreateFile creates a file in the desired path with the desired name
func (h *helperStruct) CreateFile(filePath string) (*os.File, error) {
	return os.Create(filePath)
}

// ReadFile reads the desired file
func (h *helperStruct) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// WriteFile write the content in the file
func (h *helperStruct) WriteFile(file *os.File, content string) (n int, err error) {
	return file.WriteString(content)
}
