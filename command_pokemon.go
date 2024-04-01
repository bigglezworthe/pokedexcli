package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("No Pokemon provided")
    }

    pokemonName := args[0]
    
    _ , ok := cfg.caughtPokemon[pokemonName]
    if ok {
        return errors.New("Pokemon has already been caught")
    }

    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
        return err
    }

    //arbitrarily chosen 
    const threshold = 50
    randNum :=  rand.Intn(pokemon.BaseExperience)
    fmt.Println(pokemon.BaseExperience, randNum, threshold)

    if randNum > threshold {
        return  fmt.Errorf("Failed to catch %s", pokemonName)
    }

    cfg.caughtPokemon[pokemonName] = pokemon
    fmt.Printf("You caught a %s!\n", pokemonName)
    return nil
}

func callbackInspect(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("No Pokemon provided")
    }

    pokemonName := args[0]
    
    pokemon, ok := cfg.caughtPokemon[pokemonName]
    if !ok {
        return errors.New("Only caught Pokemon may be inspected")
    }
    
    //Hardcoded for now
    fmt.Printf(" - Name: %s\n", pokemon.Name)
    fmt.Printf(" - Height: %d\n", pokemon.Height)
    fmt.Printf(" - Weight: %d\n", pokemon.Weight)
    fmt.Printf(" - Stats:\n")
    for _, stat := range pokemon.Stats {
        fmt.Printf("    - %s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Printf(" - Types:\n")
    for _, typ := range pokemon.Types {
        fmt.Printf("    - %s\n", typ.Type.Name)
    }
    
    return nil
}

//Lists caught Pokemon 
func callbackPokedex(cfg *config, args ...string) error {
    fmt.Printf("You've caught %d Pokemon:\n", len(cfg.caughtPokemon))
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    return nil

}
