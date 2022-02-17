package domain

import (
	"github.com/brisanet/cliente/enums"
	"github.com/jinzhu/gorm"
)

type Cliente struct {
	gorm.Model
	Name        string `json:"Nome" binding:"required,min=3,max=70"`
	Email       string `json:"e-mail" binding:"required,email"`
	Fone        int    `json:"fone"`
	TipoCliente enums.TipoCliente
}
