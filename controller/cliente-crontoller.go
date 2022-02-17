package controller

import (
	"net/http"

	"github.com/brisanet/cliente/domain"
	"github.com/brisanet/cliente/service"
	"github.com/gin-gonic/gin"
)

type ClienteController interface {
	FindAll() []domain.Cliente
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
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

func (c *controller) Save(ctx *gin.Context) error{
	var cliente domain.Cliente
	err := ctx.ShouldBindJSON(&cliente)
	if err!= nil {
		return err
	}
	c.service.Save(cliente)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context){
	clientes := c.service.FindAll()
	data := gin.H{
		"title": "Name Cliente",
		"clientes": clientes,
	}

	ctx.HTML(http.StatusOK,"index.html",data)


}
