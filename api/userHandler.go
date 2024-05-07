package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// GetUsers é um handler que responde com uma lista de usuários
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
}

// CreateUser é um handler para criar um novo usuário
func CreateUser(c *gin.Context) {
    // Implementação do código para criar um usuário
}
