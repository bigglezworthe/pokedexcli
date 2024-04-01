package main

import (
	"fmt"
)

//Helper function for MapNext/MapPrev
func callbackMap(cfg *config, locationURL *string) error {
    resp, err := cfg.pokeapiClient.ListLocationAreas(locationURL)
    if err != nil {
        return err
    }

    fmt.Println("Location Areas: ")
    for _, area := range resp.Results {
        fmt.Println(" - ", area.Name)
    }

    cfg.nextLocationURL = resp.Next
    cfg.prevLocationURL = resp.Previous
    
    return nil
}

func callbackMapNext(cfg *config, args ...string) error {
    return callbackMap(cfg, cfg.nextLocationURL)
}

func callbackMapPrev(cfg *config, args ...string) error {
    return callbackMap(cfg, cfg.prevLocationURL)
}

