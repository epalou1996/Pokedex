package main

import (
	"errors"
	"fmt"
)

// Struct que le daremos a la locacion especifica para buscar asi los pokemons que se hayan en la region especifica

type DetailedLocation struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, args ...string) error {
	// Verificamos que tengamos el argumento
	if len(args) < 1 {
		return errors.New("necesitamos definir un lugar que explorar")
	}

	// Inicializamos la estructuracion de la respuesta y la url en la que buscaremos la info
	var response DetailedLocation
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	err := FetchData(&url, &response)
	if err != nil {
		return err
	}

	for _, pokemon := range response.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
