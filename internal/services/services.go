package services

import (
	"github.com/vix993/shakemon/domain/entities"
	"github.com/vix993/shakemon/internal/api/models"
)

//go:generate mockgen -source services.go -destination mock/services.go -package=mock
type TranslateService interface {
	GetShakespeareTranslation(request *models.ShakespeareTranslationRequest) (res *models.ShakespeareTranslationResponse, err error)
	GetPokemon(name string) (res *entities.Pokemon, err error)
	DoGetPokemonThenTranslate(name string) (res *entities.Pokemon, err error)
}

type TranslateServiceClient struct {
	TranslateService
}
