package main

import (
 "pkdexcli/internal/pokeapi"
 "fmt"
 "errors"
)

func commandExplore(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("Name didn't macth")
  }

  name := args[0]

  explores, err := pokeapi.GetExploration(name)
  if err != nil {
    return err
  }
 fmt.Printf("Exploring %s...\n", explores.Name)
 fmt.Println("Found Pokemon:")
 for _, encounter := range explores.PokemonEncounters {
    fmt.Printf(" - %s\n",encounter.Pokemon.Name)
 }

  return nil
}

