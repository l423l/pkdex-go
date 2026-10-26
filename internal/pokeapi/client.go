package pokeapi

import (
 "pkdexcli/internal/pokecache"
 "time"
 "io"
 "net/http"
)

var cache = pokecache.NewCache(5 * time.Minute)

func locationsData(url string) ([]byte, error) {
 data, ok := cache.Get(url)
 if ok {
   return data, nil
 }

 locations, err := http.Get(url)
 if err != nil {
   return nil, err
 }
 defer locations.Body.Close()

 location, err := io.ReadAll(locations.Body)
 if err != nil {
   return nil, err
 }

 cache.Add(url, location)

 return location, nil
}
