package main

import (
 "pkdexcli/internal/pokeapi"
 "fmt"
)

func commandMapb(cfg *config, args ...string) error {
 previous := ""

 if cfg.previous != nil {
   previous = *cfg.previous
 }

 locations, err := pokeapi.GetLocationAreas(previous)
 if err != nil {
   return err
 }

 for _, location := range locations.Results {
   fmt.Println(location.Name)
 }
 cfg.previous = locations.Previous

 return nil
}
