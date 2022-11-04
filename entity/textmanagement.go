package entity

type TextManagement struct {
	ID                 uint64 `json:"id" gorm:"primaryKey;column:id"`
	TextData           string `json:"text_data" binding:"required" gorm:"type:string;column:text_data;not null"`
	Encryption         *bool  `json:"encryption" binding:"required" gorm:"type:boolean;column:encryption;not null"`
	KeySize            uint64 `json:"key_size" binding:"oneof=1024 2048 4096" gorm:"type:uint;column:key_size"`
	Uuid               string `json:"uuid" gorm:"unique;type:string;size:36;column:uuid;not null"`
	PrivateKeyPassword string `json:"private_key_password" gorm:"type:string;column:private_key_password"`
	PrivateKey         string `json:"private_key" gorm:"type:string;column:private_key"`
}
