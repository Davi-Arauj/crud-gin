package controller

import (
	"github.com/brisanet/cliente/domain"
	"github.com/brisanet/cliente/service"
	"github.com/gin-gonic/gin"
)

type ClienteController interface {
	FindAll() []domain.Cliente
	Save(ctx *gin.Context) domain.Cliente
}

type controller struct {
	service service.ClienteService
}

func New(service service.ClienteService) ClienteController {
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []domain.Cliente {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) domain.Cliente {
	var cliente domain.Cliente
	ctx.BindJSON(&cliente)
	c.service.Save(cliente)
	return cliente
}
