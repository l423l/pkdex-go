package pokeapi


type Locations struct {
 Count    int    `json:"count"`
 Next     string `json:"next"`
 Previous *string    `json:"previous"`
 Results  []struct {
   Name string `json:"name"`
   URL  string `json:"url"`
 } `json:"results"`
}

type ExploreLocation struct {
 Name string `json:"name"`
 PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type Pokemon struct {
 Name string `json:"name"`
}

type PokemonEncounters struct {
 Pokemon Pokemon `json:"pokemon"`
}
