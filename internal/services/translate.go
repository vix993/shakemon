package services

import (
	"github.com/vix993/shakemon/domain/entities"
	"github.com/vix993/shakemon/internal/api/models"
)

func (p *TranslationService) DoTranslate(name string) (res *entities.Pokemon, err error) {
	res, err = p.GetPokemon(name)
	if err != nil {
		return
	}

	translation, err := p.GetShakespeareTranslation(&models.ShakespeareTranslationRequest{
		Text: res.Description,
	})
	if err != nil {
		return
	}

	res.Description = translation.Contents.Translated
	return
}
