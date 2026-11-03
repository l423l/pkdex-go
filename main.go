package main

import (
 "pkdexcli/internal/pokeapi"
)

func main() {
 cfg := &config{
  commands: getCommands(),
  pokedex: make(map[string]pokeapi.Pokemon),
 }
 startRepl(cfg)
}
