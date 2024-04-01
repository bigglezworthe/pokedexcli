package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
    endpoint := "/pokemon/" + pokemonName
    pokemonURL := base_url + endpoint 

    //Check cache
    dat, ok := c.cache.Get(pokemonURL)

    if !ok {
        req, err := http.NewRequest("GET", pokemonURL, nil)
        if err != nil{
            return Pokemon{}, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return Pokemon{}, err
        }
        defer resp.Body.Close()

        if resp.StatusCode > 399 {
            return Pokemon{},
            fmt.Errorf("bad status code: %v", resp.StatusCode)
        }

        dat, err = io.ReadAll(resp.Body)
        if err != nil {
            return Pokemon{}, err
        }

        c.cache.Add(pokemonURL, dat)
    }

    pokemon := Pokemon{}
    err := json.Unmarshal(dat, &pokemon)
    if err != nil {
        return Pokemon{}, err
    }

    return pokemon, nil
}
