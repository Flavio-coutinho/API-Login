// handlers/auth.go
package handlers

import (
	"github.com/Flavio-coutinho/API-Login/config"
	"github.com/Flavio-coutinho/API-Login/models"
	"github.com/Flavio-coutinho/API-Login/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(c *gin.Context) {
    var user models.User
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(400, gin.H{"error": "Dados inválidos"})
        return
    }

    db := config.SetupDB()


    // Criar hash da senha antes de salvar
    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    db.Create(&user)
    c.JSON(200, gin.H{"message": "Usuário registrado com sucesso"})
}

func Login(c *gin.Context) {
    var user models.User
    err := c.ShouldBindJSON(&user)
    if err != nil {
        c.JSON(400, gin.H{"error": "Dados inválidos"})
        return
    }

    db := config.SetupDB()


    var existingUser models.User
    if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(404, gin.H{"error": "Usuário não encontrado"})
        } else {
            c.JSON(500, gin.H{"error": "Erro no servidor"})
        }
        return
    }

    if !utils.CheckPasswordHash(user.Password, existingUser.Password) {
        c.JSON(401, gin.H{"error": "Senha incorreta"})
        return
    }

    token := utils.GenerateToken(existingUser.ID)
    c.JSON(200, gin.H{"token": token})
}
