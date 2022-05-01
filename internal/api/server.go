package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vix993/shakemon/internal/api/controllers"
	"github.com/vix993/shakemon/internal/config"
	"github.com/vix993/shakemon/internal/services"
)

func Run(ctx context.Context) {
	router := gin.Default()

	router.GET("/health-check", healthCheck)

	serverConfig := cors.DefaultConfig()

	router.Use(cors.New(serverConfig))

	shakemonController := controllers.NewShakemonController(services.NewTranslationService(http.Client{}))
	shakemonRouter := router.Group("/pokemon")
	{
		shakemonRouter.GET("/:name", shakemonController.Get)
	}

	port := config.Get().Port
	if port == "" {
		port = ":8080"
	}

	log.Printf("listening on port %s", port)
	router.Run(port)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server is running",
	})
}
