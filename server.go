package main

import (
	"io"
	"os"

	"github.com/brisanet/cliente/controller"
	"github.com/brisanet/cliente/middlewares"
	"github.com/brisanet/cliente/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	clienteService service.ClienteService = service.New()
	clienteController controller.ClienteController = controller.New(clienteService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())
		

	server.GET("/clientes", func (ctx *gin.Context)  {
		ctx.JSON(200, clienteController.FindAll())
	})

	server.POST("/clientes", func (ctx * gin.Context)  {
		ctx.JSON(200,clienteController.Save(ctx))
	})

	server.Run(":8080")

}
