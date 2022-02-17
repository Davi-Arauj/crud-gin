package service

import (
	"github.com/brisanet/cliente/db"
	"github.com/brisanet/cliente/domain"
)

type ClienteService interface {
	Save(domain.Cliente) domain.Cliente
	FindAll() []domain.Cliente
	FindById(int) domain.Cliente
}

type clienteService struct {
	clientes []domain.Cliente
}

func New() ClienteService {
	return &clienteService{}
}

func (service *clienteService) Save(cliente domain.Cliente) domain.Cliente {
	db.DB.Create(&cliente)
	return cliente
}
func (service *clienteService) FindAll() []domain.Cliente {
	
	return service.clientes
}

func (service *clienteService) FindById(id int) domain.Cliente {
	clientes := service.FindAll()
	cli := domain.Cliente{}
	for indice, v := range clientes {
		if indice == id {
			cli = v
		} else {

		}
	}

	return cli
}
