package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
    endpoint := "/location-area?offset=0&limit=20" //Include full URL for caching
    locationURL := base_url + endpoint
    
    //Default to first page if something goes wrong
    if pageURL != nil {
        locationURL = *pageURL
    }

    //Check cache
    dat, ok := c.cache.Get(locationURL)
    
    //Fetch data on cache miss
    if !ok {             
        //Attempt to make request
        req, err := http.NewRequest("GET", locationURL, nil)
        if err != nil{
            return LocationAreaResp{}, err
        }

        //Do request
        resp, err := c.httpClient.Do(req)
        if err != nil{
            return LocationAreaResp{}, err
        }
        defer resp.Body.Close()  //close at end of function 

        //Get status code (anything below 400 should be handled by server)
        if resp.StatusCode > 399 {
            return LocationAreaResp{}, 
                fmt.Errorf("bad status code: %v", resp.StatusCode)
        }

        //read data as bytes 
        dat, err = io.ReadAll(resp.Body)
        if err != nil {        
            return LocationAreaResp{}, err
        }

        //cache data 
        c.cache.Add(locationURL, dat)
    }

    //convert bytes to proper structure
    locationAreasResp := LocationAreaResp{}
    err := json.Unmarshal(dat, &locationAreasResp)
    if err != nil{
        return LocationAreaResp{}, err
    }

    return locationAreasResp, nil
}

func (c *Client) ListLocationArea(locationAreaName string) (LocationArea, error) {
    endpoint := "/location-area/" + locationAreaName
    locationURL := base_url + endpoint
    
    //Check cache
    dat, ok := c.cache.Get(locationURL)
    
    //Fetch data on cache miss
    if !ok {             
        //Attempt to make request
        req, err := http.NewRequest("GET", locationURL, nil)
        if err != nil{
            return LocationArea{}, err
        }

        //Do request
        resp, err := c.httpClient.Do(req)
        if err != nil{
            return LocationArea{}, err
        }
        defer resp.Body.Close()  //close at end of function 

        //Get status code (anything below 400 should be handled by server)
        if resp.StatusCode > 399 {
            return LocationArea{}, 
                fmt.Errorf("bad status code: %v", resp.StatusCode)
        }

        //read data as bytes 
        dat, err = io.ReadAll(resp.Body)
        if err != nil {        
            return LocationArea{}, err
        }

        //cache data 
        c.cache.Add(locationURL, dat)
    }

    //convert bytes to proper structure
    locationArea := LocationArea{}
    err := json.Unmarshal(dat, &locationArea)
    if err != nil{
        return LocationArea{}, err
    }

    return locationArea, nil
}
