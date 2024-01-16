package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartRepl inicia el REPL con una configuracion dada. esta nos sirve para acceder a distintos elementos segun las necesidades del usuario.
func StartRepl(c *config) {

	// Mensaje inicial del REPL
	fmt.Println("Hola soy Rotom, talvez me conozcas con otra forma, pero ahora puedo darte mucha informacion del mundo de Pokemon, en que puedo ayudarte?")
	reader := bufio.NewScanner(os.Stdin)
	for {
		// Recolectamos la informacion que el usuario ingresa.
		fmt.Print(" Pokedex > ")
		reader.Scan()

		// Inicializamos cleanInputs, nos devuelve un array de strings, de los cuales tomaremos el primero dentro de la variable comando
		words := cleanInput(reader.Text())
		comando := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		fmt.Println(args)

		// de existir este comando en getCommands, nos devolvera la estructura asociada a este en la variable command, y un true, de lo contrario un nil y false
		// los continues permiten que el programa siga abierto.
		command, exists := getCommands()[comando]

		if exists {
			err := command.callback(c, args...)
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

// Funcion sencilla con modulo strings, que nos sirve para modificar un poco el texto
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// Los structs que nos serviran para establecer las distintas rutas que debe hacer la pokedex para conectarse con la API y tener informacion.
type config struct {
	Location MapsURLs
}
type MapsURLs struct {
	Next     *string
	Previous *string
}

// El struct que define que elementos tiene cada comando, su nombre y descripcion nos ayudan para el comando help
// y el callback sera la funcion asociada al comando.
type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	// Establecemos los comandos en un mapa como fue sugerido. sera un mapa de "comando" -> struct(comando)
	// La funcion callback se establece sin parentesis para luego establecerse mas tarde
	commandsMap := make(map[string]cliCommand)
	commandsMap["help"] = cliCommand{"help", "Con este comando podre darte ayuda", commandHelp}
	commandsMap["exit"] = cliCommand{"exit", "Con este comando cerraras la Pokedex", commandExit}
	commandsMap["map"] = cliCommand{"map", "Para ver los siguientes 20 lugares que se encuentran en este mundo.", commandMap}
	commandsMap["mapb"] = cliCommand{"mapb", "Para ver los anteriores 20 lugares que se encuentran en este mundo.", commandMapb}
	commandsMap["explore"] = cliCommand{"explore", "Veras los pokemon que encuentras en la zona que busques.", commandExplore}

	return commandsMap

}
