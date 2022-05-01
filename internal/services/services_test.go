package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vix993/shakemon/domain/entities"
	"github.com/vix993/shakemon/internal/api/models"
	"github.com/vix993/shakemon/internal/services/mock"
)

func TestDoGetPokemonThenTranslate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTranslateService := mock.NewMockTranslateService(controller)

	pokemonExample := entities.Pokemon{
		Name:        "charizard",
		Description: "some shakesperean description",
		SpriteUrl:   "some url",
	}

	mockTranslateService.EXPECT().
		DoGetPokemonThenTranslate("charizard").
		Return(&pokemonExample, nil)

	p := TranslateServiceClient{
		TranslateService: mockTranslateService,
	}

	res, err := p.DoGetPokemonThenTranslate("charizard")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.IsType(t, &entities.Pokemon{}, res)
}

func TestGetPokemon(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTranslateService := mock.NewMockTranslateService(controller)

	pokemonExample := entities.Pokemon{
		Name:        "charizard",
		Description: "some shakesperean description",
		SpriteUrl:   "some url",
	}

	mockTranslateService.EXPECT().
		GetPokemon("charizard").
		Return(&pokemonExample, nil)

	p := TranslateServiceClient{
		TranslateService: mockTranslateService,
	}

	res, err := p.GetPokemon("charizard")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.IsType(t, &entities.Pokemon{}, res)
}

func TestGetShakespeareTranslation(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockTranslateService := mock.NewMockTranslateService(controller)

	pokemonExample := models.ShakespeareTranslationResponse{}

	mockTranslateService.EXPECT().
		GetShakespeareTranslation(&models.ShakespeareTranslationRequest{Text: "hello"}).
		Return(&pokemonExample, nil)

	p := TranslateServiceClient{
		TranslateService: mockTranslateService,
	}

	res, err := p.GetShakespeareTranslation(&models.ShakespeareTranslationRequest{Text: "hello"})

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.IsType(t, &models.ShakespeareTranslationResponse{}, res)
}
