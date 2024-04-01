package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("No location area provided")
    }

    locationAreaName := args[0]

    locationArea, err := cfg.pokeapiClient.ListLocationArea(locationAreaName)
    if err != nil {
        return err
    }

    fmt.Printf("Pokemon in %s:\n", locationArea.Name)

    for _, encounter := range locationArea.PokemonEncounters {
        fmt.Println(" - ", encounter.Pokemon.Name)
    }

    return nil

}

