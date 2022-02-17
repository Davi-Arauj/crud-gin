package main

import (
	"net/http"

	"github.com/brisanet/cliente/controller"

	"github.com/brisanet/cliente/middlewares"
	"github.com/brisanet/cliente/repository"
	"github.com/brisanet/cliente/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	clienteRepository repository.ClienteRepository = repository.NewClienteRepository()

	clienteService service.ClienteService = service.New(clienteRepository)

	clienteController controller.ClienteController = controller.New(clienteService)
)

//https://gitlab.com/pragmaticreviews/golang-gin-poc/-/blob/gorm/server.go
func main() {

	defer clienteRepository.CloseDB()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/cliente", func(ctx *gin.Context) {
			ctx.JSON(200, clienteController.FindAll())
		})

		apiRoutes.POST("/cliente", func(ctx *gin.Context) {
			err := clienteController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Cliente Inserido com sucesso!  :) "})
			}
		})

	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/clientes", clienteController.ShowAll)
	}

	server.Run(":8080")

}
