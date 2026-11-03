package main

import (
 "pkdexcli/internal/pokeapi"
 "fmt"
 "errors"
 "math/rand"
)

func commandCatch(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("Name didn't macth")
  }

  name := args[0]

  explores, err := pokeapi.GetPokemon(name)
  if err != nil {
    return err
  }
 fmt.Printf("Throwing a Pokeball at %s...\n", explores.Name)
 limit := explores.BaseExperience
 if limit <= 0 {
   limit = 1
 }

 if rand.Intn(limit) < 56  {
   cfg.pokedex[explores.Name] = explores
   fmt.Printf("%s was caught!\n", explores.Name)
 } else {
   fmt.Printf("%s escaped!\n", explores.Name)
 }

  return nil
}
