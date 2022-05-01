package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vix993/shakemon/domain/entities"
)

type ShakemonServiceClient interface {
	DoGetPokemonThenTranslate(name string) (res *entities.Pokemon, err error)
}

type ShakemonController struct {
	client ShakemonServiceClient
}

func NewShakemonController(shakemonServiceClient ShakemonServiceClient) *ShakemonController {
	return &ShakemonController{
		client: shakemonServiceClient,
	}
}

func (p *ShakemonController) Get(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "pokemon name is required",
		})
		return
	}

	res, err := p.client.DoGetPokemonThenTranslate(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
