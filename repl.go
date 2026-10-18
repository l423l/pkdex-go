package main

import (
 "strings"
 "bufio"
 "os"
 "fmt"
)

func cleanInput(text string) []string {
 lowerCase := strings.ToLower(text)
 splitText := strings.Fields(lowerCase)
 return splitText
}

func startRepl() {
 scanner := bufio.NewScanner(os.Stdin)
 for {
  fmt.Print("Pokedex > ")
  scanner.Scan()
  text := scanner.Text()

  words := cleanInput(text)

  firstWord := words[0]

  fmt.Printf("Your command was: %s\n", firstWord)
 }
}
