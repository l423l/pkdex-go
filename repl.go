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
 callback func(*config, ...string) error
}

type config struct {
 commands map[string]cliCommand
 next *string
 previous *string
}

func cleanInput(text string) []string {
 lowerCase := strings.ToLower(text)
 splitText := strings.Fields(lowerCase)
 return splitText
}

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
  "map": {
    name: "map",
    description: "Displays the locations",
    callback: commandMap,
  },
  "mapb": {
    name: "mapb",
    description: "Displays the locations",
    callback: commandMapb,
  },
  "explore": {
    name: "explore",
    description: "Gotta catch'em all",
    callback: commandExplore,
  },
 }
}

func startRepl(cfg *config) {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    text := scanner.Text()

    words := cleanInput(text)
    if len(words) == 0 {
      continue
    }

    firstWord := words[0]
    args := words[1:]

    command, ok := cfg.commands[firstWord]
    if !ok {
      fmt.Print("Unknown command")
    } else {
        err := command.callback(cfg, args...)
      if err != nil {
          fmt.Print(err)
      }
    }
  }
}
