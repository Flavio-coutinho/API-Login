package routes

import (
	"github.com/Flavio-coutinho/API-Login/handlers"
	"github.com/Flavio-coutinho/API-Login/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.POST("login", handlers.Login)

	protected := router.Group("/protected")
    protected.Use(utils.AuthMiddleware()) // Aplica o middleware de autenticação
    {
        protected.GET("/profile", handlers.Profile)
    }

	return router
}