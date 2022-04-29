package services

import (
	"fmt"
	"math/rand"
	"time"

	pokeApiGo "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/vix993/shakemon/domain/entities"
)

func (p *TranslationService) GetPokemon(name string) (res *entities.Pokemon, err error) {
	response, err := pokeApiGo.Pokemon(name)
	if err != nil {
		return
	}

	res = makePokemonFromPokeApiResponse(response)
	return
}

func makePokemonFromPokeApiResponse(pokemon structs.Pokemon) *entities.Pokemon {
	pokemonType := pokemon.Types[0].Type.Name
	availableMoves := len(pokemon.Moves)
	moveIndex1, moveIndex2 := randomizeMoves(availableMoves)
	move1 := pokemon.Moves[moveIndex1].Move.Name
	move2 := pokemon.Moves[moveIndex2].Move.Name
	return &entities.Pokemon{
		Name: pokemon.Name,
		Description: fmt.Sprintf(
			"%s is a very strong %s type pokemon. It is number %d in the pokedex and has an average weight of %d. This reinforcement to your team can learn attacks like %s and %s. You have to catch them all!",
			pokemon.Name,
			pokemonType,
			pokemon.ID,
			pokemon.Weight,
			move1,
			move2,
		),
		SpriteUrl: pokemon.Sprites.FrontDefault,
	}
}

func randomizeMoves(len int) (i1 int, i2 int) {
	rand.Seed(time.Now().UnixNano())
	i1 = rand.Intn(len)
	i2 = rand.Intn(len)
	for i1 == i2 {
		i2 = rand.Intn(len)
	}
	return
}
