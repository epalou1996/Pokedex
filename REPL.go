package main

import (
	"fmt"
	"strings"
)

func StartRepl() {

	for {
		var mensaje string
		fmt.Println("Hola Bienvenido a la Pokedex, en que puedo ayudarte?")
		fmt.Print(" Pokedex > ")
		fmt.Scanln(&mensaje)
		words := cleanInputs(mensaje)
		comando := words[0]

		command, exists := getCommands()[comando]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Comando desconocido")
			continue
		}
	}
}

func cleanInputs(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	// Establecemos los comandos en un mapa como fue sugerido. sera un mapa de "comando" -> struct(comando)
	// La funcion callback se establece sin parentesis para luego establecerse mas tarde
	commandMap := make(map[string]cliCommand)
	commandMap["help"] = cliCommand{"help", "Pronto podre ayudarte!", commandHelp}
	commandMap["exit"] = cliCommand{"exit", "Hasta luego!", commandExit}

	return commandMap

}
