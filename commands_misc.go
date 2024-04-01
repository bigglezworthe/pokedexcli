package main

import (
    "fmt"
    "os"
)

func callbackHelp(cfg *config, args ...string) error {
    fmt.Println("welcome to the pokedex help menu!")
    fmt.Println("Here are your available commands:" )

    commands := getCommands()

    for _, command := range commands {
        fmt.Printf(" - %s: %s \n", 
            command.name, 
            command.description,
        )
    }

    fmt.Println("")
    return nil
}

func callbackExit(cfg *config, args ...string) error {
    os.Exit(0)
    return nil
}
