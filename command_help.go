package main

import "fmt"

// Es una funcion simple que nos permite mostrar todos los comandos que hay en nuestra pokedex
func commandHelp(c *config, args ...string) error {
	fmt.Println("Uso:")
	fmt.Println()
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v", command.name, command.description)
		fmt.Println()
	}
	fmt.Println()
	return nil

}
