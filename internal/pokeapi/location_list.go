package pokeapi

import (
 "encoding/json"
)

func locationsJSON(data []byte) (Locations,error) {
 var locations Locations

 err := json.Unmarshal(data, &locations)
 if err != nil {
   return Locations{}, err
 }

 return locations, nil
}

func GetLocationAreas(url string) (Locations, error) {
 if url == "" {
   url = baseURL
 }

 data, err := locationsData(url)
 if err != nil {
   return Locations{}, err
 }

 locations, err := locationsJSON(data)
 if err != nil {
   return Locations{}, err
 }

 return locations, nil
}
