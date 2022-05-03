# Shakémon API:
## *If Shakespeare wrote the Pokédex*

## Content

1. [Description](#Motivation) 
2. [Requirements](#Requirements)
3. [API](#API)
    - [Pokémon](#Pokémon)
    - [HealthCheck](#HealthCheck)

## Status

![Tests](https://github.com/vix993/shakemon/actions/workflows/test.yml/badge.svg)

## Description

I was tasked with creating an application that would return a description of a pokemon in the style of William Shakespeare.

The result of this prompt was [`Shakémon`](https://shakemon.vercel.app/). This repository contains the backend for this app that has been deployed to GCP Cloud Run.

This application currently provides two resources, the `/health-check` endpoint - that informs whether the server is running, and the `/pokemon/{name}` endpoint - that executes the search and translation of a given pokemon.

The repository containing the front end of this app can be acessed here: https://github.com/vix993/shakemon-front

You can run the backend:
- `git clone git@github.com:vix993/shakemon.git shakemon-api`
- `cd shakemon-api`
- `docker build . -t shakemon_api`
- `docker run -p 8080:8080 -d shakemon_api`

For Cross Origin Sharing in local environment:
- `docker build -f Dockerfile.nginx -t cors .`
- `docker run cors`

Then the frontend:
- `git clone git@github.com:vix993/shakemon-front.git ../shakemon-front`
- `cd ../shakemon-front`
- `docker build -t shakemon .`
- `docker run -p 3000:3000 shakemon`

Without docker:
- [`install golang`](https://go.dev/doc/install)
- `go mod tidy`
- `go run cmd/main.go`

The unit tests are in `internal/services/services_test.go` and can be run like so:
- `make test`
- alternatively you can read the [ci logs](https://github.com/vix993/shakemon/actions) for the output.

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
```{
    "name": "charizard",
    "description": "Charizard is a very stout fire type pokemon. 't is number 6 in the pokedex and hath an average weight of 905. 't is a valorous reinforcement to thy team and is able to learneth attacks like bodkins-dance,  scary-visage and facade. Thy opponents wilt beest regretting facing off with thee. Thee hath't to catcheth those folk all!",
    "spriteUrl": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/6.png"
}```
## *HealthCheck*

#### GET
Get server status.
#### Endpoint
`GET /health-check`
#### Response parameter
##### 200
```{
    "message": "server is running",
}```
