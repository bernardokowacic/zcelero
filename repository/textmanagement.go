package repository

type TextManagementInterface interface {
	Get()
}

type textManagementRepositoryStruct struct {
}

func NewRepository() TextManagementInterface {
	return &textManagementRepositoryStruct{}
}

func (p *textManagementRepositoryStruct) Get() {
}
