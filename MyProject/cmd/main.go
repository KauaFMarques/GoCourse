package main

import (
	"meu-projeto/controller"
	"meu-projeto/db"
	"meu-projeto/repository" // ADICIONE ISSO
	"meu-projeto/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	// CORREÇÃO: Removida a chave '{' que estava sobrando aqui
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	productRepository := repository.NewProductRepository(dbConnection)

	// Camada de usecase
	productUseCase := usecase.NewProductUsecase(productRepository)

	// Camada de controllers
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductsById)

	server.Run(":8000")
}
