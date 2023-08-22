package handlers

import (
    "github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
    claims, exists := c.Get("claims")
    if !exists {
        c.JSON(401, gin.H{"error": "Token inválido"})
        return
    }

    userID := uint(claims.(map[string]interface{})["sub"].(float64))

    // Aqui você pode consultar o banco de dados ou outra fonte para obter as informações do usuário usando o userID

    // Simplesmente retornando as informações do usuário neste exemplo
    user := map[string]interface{}{
        "id":    userID,
        "name":  "John Doe",
        "email": "john@example.com",
    }

    c.JSON(200, gin.H{"user": user})
}
