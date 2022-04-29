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

	shakemonController := controllers.NewShakemonController(services.NewTranslationService(http.Client{}))
	shakemonRouter := router.Group("/pokemon")
	{
		shakemonRouter.GET("/:name", shakemonController.Get)
	}

	router.Run(":8080")
}
