package main

import (
	"io"
	"net/http"
	"os"

	"github.com/brisanet/cliente/controller"
	"github.com/brisanet/cliente/db"
	"github.com/brisanet/cliente/domain"
	"github.com/brisanet/cliente/middlewares"
	"github.com/brisanet/cliente/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	clienteService    service.ClienteService       = service.New()
	clienteController controller.ClienteController = controller.New(clienteService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	db.ConnectDB()

	db.DB.AutoMigrate(&domain.Cliente{})

	setupLogOutput()

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
		apiRoutes.GET("/cliente/:id", func(ctx *gin.Context) {
			ctx.JSON(200, clienteController.FindById)
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

	defer db.DB.Close()
}
