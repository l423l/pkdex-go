package main

import (
 "pkdexcli/internal/pokeapi"
 "fmt"
)

func fecthlocations(url string) (pokeapi.Locations, error) {
 locations, err := pokeapi.GetLocationAreas(url)
 if err != nil {
   return pokeapi.Locations{}, err
 }

 return locations, nil
}

func commandMap(cfg *config, args ...string) error {
 next := ""

 if cfg.next != nil {
   next = *cfg.next
 }

 locations, err := pokeapi.GetLocationAreas(next)
 if err != nil {
   return err
 }

 for _, location := range locations.Results {
   fmt.Println(location.Name)
 }
 cfg.next = &locations.Next

 return nil
}
