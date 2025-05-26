package internal

import (
	"encoding/json"
	"net/http"
)

func GetLocations(url string) (data LocationAreas, err error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return LocationAreas{}, err
	}

	return data, nil
}

func GetEncounteredPokemon(url string) (data AreaData, err error) {
	res, err := http.Get(url)
	if err != nil {
		return AreaData{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return AreaData{}, err
	}

	return data, nil
}

func GetPokemonData(url string) (data PokemonData, err error) {
	res, err := http.Get(url)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return PokemonData{}, err
	}

	return data, nil
}
