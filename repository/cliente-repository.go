package repository

import (
	"fmt"
	"os"

	"github.com/brisanet/cliente/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ClienteRepository interface {
	Save(cliente domain.Cliente)
	Update(cliente domain.Cliente)
	Delete(cliente domain.Cliente)
	FindAll() []domain.Cliente
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewClienteRepository() ClienteRepository {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	db, err := gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to database successfully")
	}

	db.AutoMigrate(&domain.Cliente{})

	return &database{
		connection: db,
	}
}
func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(cliente domain.Cliente) {
	db.connection.Create(&cliente)
}
func (db *database) Update(cliente domain.Cliente) {
	db.connection.Save(&cliente)
}
func (db *database) Delete(cliente domain.Cliente) {
	db.connection.Delete(&cliente)
}
func (db *database) FindAll() []domain.Cliente {
	var clientes []domain.Cliente
	db.connection.Set("gorm:auto_preload", true).Find(&clientes)
	return clientes
}
