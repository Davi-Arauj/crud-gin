package service

import (
	"fmt"

	"github.com/brisanet/cliente/domain"
	"github.com/brisanet/cliente/repository"
)

type ClienteService interface {
	Save(domain.Cliente) error
	Update(domain.Cliente) error
	Delete(domain.Cliente) error
	FindById(domain.Cliente) (domain.Cliente)
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

func (service *clienteService) FindById(cliente domain.Cliente) (domain.Cliente) {
	fmt.Println("Service->",cliente)
	return service.repository.FindById(cliente)
}
