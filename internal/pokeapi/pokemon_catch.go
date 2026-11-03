package pokeapi

import (
 "encoding/json"
)

func pokemonCatchJSON(data []byte) (Pokemon, error) {
 var pokemons Pokemon

 err := json.Unmarshal(data, &pokemons)
 if err != nil {
   return Pokemon{}, err
 }

 return pokemons, nil
}

func GetPokemon(name string) (Pokemon, error) {
 url := pokemonBaseURL + name
 data, err := locationsData(url)
 if err != nil {
   return Pokemon{}, err
 }
 pokemons, err := pokemonCatchJSON(data)
 if err != nil {
   return Pokemon{}, err
 }
 return pokemons, nil
}
