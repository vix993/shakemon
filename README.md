# Shakémon API:
## *If Shakespeare wrote the Pokédex*

## Content

1. [Description](#Motivation) 
2. [Requirements](#Requirements)
3. [API](#API)
    - [Pokémon](#Pokémon)
    - [HealthCheck](#HealthCheck)
4. [Additional Notes](#API)

## Status

![Tests](https://github.com/vix993/project-lifeline/actions/workflows/test.yml/badge.svg)  ![Health Check](https://github.com/vix993/project-lifeline/actions/workflows/health_check.yml/badge.svg)

## Description

I was tasked with creating an API that would return a description of a pokemon in the style of William Shakespeare.

What came out was [`ShakemonAPI`](https://shakemon-zu2jrydmba-lm.a.run.app).

This application currently provides two resources, the `/health-check` endpoint - that informs whether the server is running, and the `/pokemon/{name}` endpoint - that executes the search and translation of a given pokemon.

You can run it with the following commands: 
- `docker build . -t shakemon`
- `docker run -p 8080:8080 -d shakemon`

Without docker:
- [`install golang`](https://go.dev/doc/install)
- `go mod tidy`
- `go run cmd/main.go`

## Requirements

Go or docker.

## Api
## *Pokémon*

#### GET
Get the Shakespearean description of a pokémon.
#### Endpoint
`GET /pokemon/{name}`
#### Response parameter
##### 200
`{
    "name": "charizard",
    "description": "Charizard is a very stout fire type pokemon. 't is number 6 in the pokedex and hath an average weight of 905. 't is a valorous reinforcement to thy team and is able to learneth attacks like bodkins-dance,  scary-visage and facade. Thy opponents wilt beest regretting facing off with thee. Thee hath't to catcheth those folk all!",
    "spriteUrl": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/6.png"
}`
## *HealthCheck*

#### GET
Get server status.
#### Endpoint
`GET /health-check`
#### Response parameter
##### 200
`{
    "message": "server is running",
}`
