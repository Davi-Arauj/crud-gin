package domain

type Endereco struct {
	ID         int64  `json:"id" gorm:"primary_key;auto_increment"`
	Logradouro string `json:"logradouro" binding:"required,min=3,max=70"`
	Bairro     string `json:"bairro" binding:"required,min=3,max=70"`
	numero     int64  `json:"numero"`
}
