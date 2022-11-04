package entity

type TextManagement struct {
	TextData           string `json:"text_data" binding:"required"`
	Encryption         *bool  `json:"encryption" binding:"required"`
	KeySize            uint64 `json:"key_size"`
	Uuid               string `json:"uuid"`
	PrivateKeyPassword string `json:"private_key_password"`
	PrivateKey         string `json:"private_key"`
}
