package pokeapi

import (
 "encoding/json"
)

func exploreLocationJSON(data []byte) (ExploreLocation, error) {
 var locations ExploreLocation

 err := json.Unmarshal(data, &locations)
 if err != nil {
   return ExploreLocation{}, err
 }

 return locations, nil
}

func GetExploration(name string) (ExploreLocation, error) {
 url := baseURL + name
 data, err := locationsData(url)
 if err != nil {
   return ExploreLocation{}, err
 }
 locations, err := exploreLocationJSON(data)
 if err != nil {
   return ExploreLocation{}, err
 }
 return locations, nil
}
