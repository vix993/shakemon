package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vix993/shakemon/internal/api/controllers"
	"github.com/vix993/shakemon/internal/services"
)

func Run(ctx context.Context) {
	router := gin.Default()

	router.GET("/health-check", healthCheck)

	shakemonController := controllers.NewShakemonController(services.NewTranslationService(http.Client{}))
	shakemonRouter := router.Group("/pokemon")
	{
		shakemonRouter.GET("/:name", shakemonController.Get)
	}

	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server is running",
	})
}
