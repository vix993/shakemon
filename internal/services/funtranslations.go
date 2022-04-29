package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/vix993/shakemon/internal/api/models"
	"github.com/vix993/shakemon/internal/config"
)

type TranslationService struct {
	client http.Client
}

func NewTranslationService(client http.Client) *TranslationService {
	return &TranslationService{
		client: client,
	}
}

func (p *TranslationService) GetShakespeareTranslation(request *models.ShakespeareTranslationRequest) (res *models.ShakespeareTranslationResponse, err error) {
	// Create request
	apiUrl := "https://api.funtranslations.com"
	resource := "/translate/shakespeare.json"
	data := url.Values{}
	data.Set("text", request.Text)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)

	funtranslationsApiKey := config.Get().FtAPIKey

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Funtranslations-Api-Secret", funtranslationsApiKey)

	// Execute request
	response, err := p.client.Do(req)
	if err != nil {
		return res, err
	}
	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("error on get shakespeare translation: code - %v, message - %v", response.StatusCode, response.Status))
		return
	}

	// Parse response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return
	}

	return
}
