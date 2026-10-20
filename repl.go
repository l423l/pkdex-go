package main

import (
 "strings"
 "bufio"
 "os"
 "fmt"
)

type cliCommand struct {
 name string
 description string
 callback func() error
}

func cleanInput(text string) []string {
 lowerCase := strings.ToLower(text)
 splitText := strings.Fields(lowerCase)
 return splitText
}

var commands map[string]cliCommand

func getCommands() map[string]cliCommand {
 return map[string]cliCommand{
  "help": {
    name: "help",
    description: "Displays a help message",
    callback: commandHelp,
  },
  "exit": {
    name: "exit",
    description: "Exit the Pokedex",
    callback: commandExit,
  },
 }
}

func startRepl() {
 commands = getCommands()
 scanner := bufio.NewScanner(os.Stdin)

 for {
  fmt.Print("Pokedex > ")
  scanner.Scan()
  text := scanner.Text()

  words := cleanInput(text)

  firstWord := words[0]

  command, ok := commands[firstWord]
  if !ok {
    fmt.Print("Unknown command")
  } else {
    err := command.callback()
    if err != nil {
      fmt.Print(err)
    }
  }
 }
}
