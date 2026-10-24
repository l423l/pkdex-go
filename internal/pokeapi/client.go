package pokeapi

import (
 "io"
 "net/http"
)

func locationsData(url string) ([]byte, error) {
 locations, err := http.Get(url)
 if err != nil {
   return nil, err
 }
 defer locations.Body.Close()

 location, err := io.ReadAll(locations.Body)
 if err != nil {
   return nil, err
 }

 return location, nil
}
