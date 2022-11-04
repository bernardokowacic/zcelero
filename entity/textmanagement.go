package entity

type TextManagement struct {
	ID                 uint64 `json:"id"`
	TextData           string `json:"text_data" binding:"required"`
	Encryption         *bool  `json:"encryption" binding:"required"`
	KeySize            uint64 `json:"key_size" binding:"oneof=1024 2048 4096"`
	Uuid               string `json:"uuid"`
	PrivateKeyPassword string `json:"private_key_password"`
	PrivateKey         string `json:"private_key"`
}