package entity

type TextManagement struct {
	ID                 uint64 `json:"id" gorm:"primaryKey;column:id"`
	TextData           string `json:"text_data" binding:"required" gorm:"type:string;column:text_data;not null"`
	Encription         bool   `json:"encription" binding:"required" gorm:"type:boolean;column:encription;not null"`
	KeySize            uint64 `json:"key_size" gorm:"type:uint;column:key_size"`
	Uuid               string `json:"uuid" gorm:"unique;type:string;size:36;column:uuid;not null"`
	PrivateKeyPassword string `json:"private_key_password" gorm:"type:string;column:private_key_password"`
}
