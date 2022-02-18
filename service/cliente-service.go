package service

import (
	"github.com/brisanet/cliente/domain"
	"github.com/brisanet/cliente/repository"
)

type ClienteService interface {
	Save(domain.Cliente) error
	Update(domain.Cliente) error
	Delete(domain.Cliente) error
	FindAll() []domain.Cliente
}

type clienteService struct {
	repository repository.ClienteRepository
}

func New(clienteRepository repository.ClienteRepository) ClienteService {
	return &clienteService{
		repository: clienteRepository,
	}
}

func (service *clienteService) Save(cliente domain.Cliente) error {
	service.repository.Save(cliente)
	return nil
}

func (service *clienteService) Update(cliente domain.Cliente) error {
	service.repository.Update(cliente)
	return nil
}

func (service *clienteService) Delete(cliente domain.Cliente) error {
	service.repository.Delete(cliente)
	return nil
}
func (service *clienteService) FindAll() []domain.Cliente {
	return service.repository.FindAll()
}
