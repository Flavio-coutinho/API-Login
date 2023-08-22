package utils

import (
	"net/http"

	"github.com/Flavio-coutinho/API-Login/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
				if tokenString == "" {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token nao fornecido"})
					c.Abort()
					return
				}
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.GetSecretKey()), nil
				})

				if err != nil || !token.Valid {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalido"})
					c.Abort()
					return
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao obter as reivindicacoes do token"})
					c.Abort()
					return
				}

				c.Set("claims", claims)

				c.Next()
        // Se o token for válido, passe para o próximo handler
        // Caso contrário, retorne um erro não autorizado (401)
    }
}