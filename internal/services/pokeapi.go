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
	moves := makeMovesDescription(availableMoves, pokemon)

	return &entities.Pokemon{
		Name: pokemon.Name,
		Description: fmt.Sprintf(
			"%s is a very strong %s type pokemon. It is number %d in the pokedex and has an average weight of %d. %s Your opponents will be regretting facing off with you. You have to catch them all!",
			pokemon.Name,
			pokemonType,
			pokemon.ID,
			pokemon.Weight,
			moves,
		),
		SpriteUrl: pokemon.Sprites.FrontDefault,
	}
}

func randomizeMoves(len int) (i1, i2, i3 int) {
	rand.Seed(time.Now().UnixNano())
	i1 = rand.Intn(len)
	i2 = rand.Intn(len)
	i3 = rand.Intn(len)
	for i1 == i2 {
		i2 = rand.Intn(len)
	}
	for i3 == i2 || i3 == i1 {
		i3 = rand.Intn(len)
	}
	return
}

func makeMovesDescription(availableMoves int, pokemon structs.Pokemon) string {
	if availableMoves == 1 {
		return fmt.Sprintf("It is a good reinforcement to your team and is able to learn %s.", pokemon.Moves[0].Move.Name)
	} else if availableMoves == 2 {
		return fmt.Sprintf("It is a good reinforcement to your team and is able to learn attacks like %s and %s.", pokemon.Moves[0].Move.Name, pokemon.Moves[1].Move.Name)
	} else {
		moveIndex1, moveIndex2, moveIndex3 := randomizeMoves(availableMoves)
		move1 := pokemon.Moves[moveIndex1].Move.Name
		move2 := pokemon.Moves[moveIndex2].Move.Name
		move3 := pokemon.Moves[moveIndex3].Move.Name
		return fmt.Sprintf(
			"It is a good reinforcement to your team and is able to learn attacks like %s, %s and %s.",
			move1,
			move2,
			move3,
		)
	}
}
