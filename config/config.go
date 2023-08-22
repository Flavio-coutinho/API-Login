// config/config.go
package config

import (
    "os"
)

// GetSecretKey retorna a chave secreta do ambiente ou uma padrão
func GetSecretKey() string {
    secretKey := os.Getenv("SECRET_KEY")

    if secretKey == "" {
        secretKey = "paocomcha" // Substitua pela chave padrão real
    }

    return secretKey
}
