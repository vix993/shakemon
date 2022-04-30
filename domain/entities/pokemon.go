package entities

type Pokemon struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SpriteUrl   string `json:"spriteUrl"`
}
