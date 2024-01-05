package main

import (
	"fmt"
	"strings"
)

func StartRepl(c *config) {
	fmt.Println("Hola soy Rotom, talvez me conozcas con otra forma, pero ahora puedo darte mucha informacion del mundo de Pokemon, en que puedo ayudarte?")
	for {
		var mensaje string
		fmt.Print(" Pokedex > ")
		fmt.Scanln(&mensaje)
		words := cleanInputs(mensaje)
		comando := words[0]

		command, exists := getCommands()[comando]
		if exists {
			err := command.callback(c)
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

type config struct {
	Location MapsURLs
}
type MapsURLs struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	// Establecemos los comandos en un mapa como fue sugerido. sera un mapa de "comando" -> struct(comando)
	// La funcion callback se establece sin parentesis para luego establecerse mas tarde
	commandsMap := make(map[string]cliCommand)
	commandsMap["help"] = cliCommand{"help", "Pronto podre ayudarte!", commandHelp}
	commandsMap["exit"] = cliCommand{"exit", "Hasta luego!", commandExit}
	commandsMap["map"] = cliCommand{"map", "Para ver los siguientes 20 lugares que se encuentran en este mundo.", commandMap}
	commandsMap["mapb"] = cliCommand{"map", "Para ver los anteriores 20 lugares que se encuentran en este mundo.", commandMapb}

	return commandsMap

}
