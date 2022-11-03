package service

import "zcelero/repository"

type TextManagementServiceInteface interface {
	Get()
}

type TextManagementService struct {
	TextManagementRepository repository.TextManagementInterface
}

func NewService(textManagementRepository repository.TextManagementInterface) TextManagementServiceInteface {
	return &TextManagementService{TextManagementRepository: textManagementRepository}
}

func (t *TextManagementService) Get() {

}
