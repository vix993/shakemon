package models

type ShakespeareTranslationRequest struct {
	Text string
}

type ShakespeareTranslationResponse struct {
	Success  Success  `json:"success"`
	Contents Contents `json:"contents"`
}

type Success struct {
	Total int `json:"total"`
}

type Contents struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"shakespeare"`
}
