package repository

import "gorm.io/gorm"

type TextManagementInterface interface {
	Get()
}

type textManagementRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewRepository(dbConn *gorm.DB) TextManagementInterface {
	return &textManagementRepositoryStruct{DbConn: dbConn}
}

func (p *textManagementRepositoryStruct) Get() {
}
