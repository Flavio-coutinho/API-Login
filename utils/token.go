package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateToken gera um token JWT
func GenerateToken(userID uint) string {
	// Defina as reivindicações do token, como ID do usuário e expiração
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		"iat": time.Now().Unix(),
	}

	// Crie o token com as reivindicações
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assine o token usando uma chave secreta
	secretKey := []byte("sua_chave_secreta") // Substitua pela sua chave secreta real
	tokenString, _ := token.SignedString(secretKey)

	return tokenString
}
