package main

import (
	"apiWeather/api" // Importe os handlers que você definiu

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma instância do Gin
	router := gin.Default()

	// Middlewares globais
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	// Define as rotas
	router.GET("/ping", api.GetUsers)    // Suponha que `GetUsers` seja uma função em `userHandler.go`
	router.POST("/ping", api.CreateUser) // Suponha que `CreateUser` seja outra função em `userHandler.go`

	// Inicia o servidor na porta especificada
	router.Run(":80") // Escuta e serve na porta 8080
}
