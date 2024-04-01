package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)

    //var text string
    all_commands := getCommands() 

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := scanner.Text()

        cleaned := cleanInput(text)
        if len(cleaned) == 0 {
            continue
        }

        command_name := cleaned[0]
        args := []string{}
        if len(cleaned) > 0 {
            args = cleaned[1:]
        }

        command, ok := all_commands[command_name]
        if !ok {
            fmt.Println("Invalid Command")
        } else {
            err := command.callback(cfg, args...)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}

type cliCommand struct {
    name string
    description string
    callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "map": {
            name: "map",
            description: "Prints the next location areas",
            callback: callbackMapNext,
        },
        "mapb": {
            name: "mapb",
            description: "Prints the previous location areas",
            callback: callbackMapPrev,
        },
        "explore": {
            name: "explore <location_area>",
            description: "Lists the Pokemon at a location area",
            callback: callbackExplore,
        },
        "catch": {
            name: "catch <pokemon>",
            description: "Try to catch a Pokemon",
            callback: callbackCatch,
        },
        "inspect": {
            name: "inspect <pokemon>",
            description: "Display stats for a captured Pokemon",
            callback: callbackInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "Display all captured Pokemon",
            callback: callbackPokedex,
        },
        "help": {
            name: "help",
            description: "Prints the help menu",
            callback: callbackHelp,
        },
        "exit": {
            name: "exit",
            description: "Exits the program",
            callback: callbackExit,
        },
    }
}

func cleanInput(str string) []string {
    lower := strings.ToLower(str)
    words := strings.Fields(lower)
    return words
}

