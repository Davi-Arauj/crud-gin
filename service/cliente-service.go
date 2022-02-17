package service

import "github.com/brisanet/cliente/domain"

type ClienteService interface {
	Save(domain.Cliente) domain.Cliente
	FindAll() []domain.Cliente
}

type clienteService struct {
	clientes []domain.Cliente
}

func New() ClienteService {
	return &clienteService{}
}

func (service *clienteService) Save(cliente domain.Cliente) domain.Cliente {
	service.clientes = append(service.clientes, cliente)
	return cliente
}
func (service *clienteService) FindAll() []domain.Cliente {
	return service.clientes
}
