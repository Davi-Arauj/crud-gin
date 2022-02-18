package repository

import "github.com/brisanet/cliente/domain"

type EnderecoRepository interface {
	Save(endereco domain.Endereco)
	Update(endereco domain.Endereco)
	Delete(endereco domain.Endereco)
	FindById(endereco domain.Endereco) domain.Endereco
	FindAll() []domain.Endereco
	CloseDB()
}


