package domain

import (
	"github.com/brisanet/cliente/enums"
	"github.com/jinzhu/gorm"
)

type Cliente struct {
	gorm.Model
	Name string `json:"Nome"`
	Email string `json:"e-mail"`
	Fone int  `json:"fone"`
	TipoCliente enums.TipoCliente 
}
