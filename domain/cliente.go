package domain

import (
	"github.com/brisanet/cliente/enums"
)

type Cliente struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `json:"Nome" binding:"required,min=3,max=70"`
	Email       string `json:"e-mail" binding:"required,email"`
	Fone        int    `json:"fone"`
	TipoCliente enums.TipoCliente `json:"tipo"`
}
