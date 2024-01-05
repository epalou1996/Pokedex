package main

import "fmt"

// Esta sera la funcion que nos ayude a mapear correctamente hacia adelante.

func commandMap(c *config) error {
	// Declarar una variable del tipo de estructura esperado
	var response LocationResponse

	// Hacer la solicitud HTTP y llenar la estructura. Esto lo hacemos con c.Next ya que en config mandamos que este elemento tendria la URL de los locations
	// Naturalmente hay mejores formas de hacerlo, pero por ahora esto nos servira.

	err := FetchData(c.Location.Next, &response)
	if err != nil {
		return err
	}

	// Establecer los cambios en las rutas.
	c.Location.Previous = response.Previous
	c.Location.Next = response.Next
	// Acceder a los resultados
	for _, location := range response.Results {
		fmt.Println(location.Name)

	}

	return nil
}

// Esta funcion sera el opuesto a la primera, ira de hacia los 20 anteriores resultados

func commandMapb(c *config) error {

	// Primero mirar si hay resultados anteriores, con la variable que pasa la config se podria.
	if *c.Location.Previous == "" {
		fmt.Println("Este comando solo sirve para retroceder en la lista de mapas, usa map para ver la lista de los distintos mapas")
		return nil
	}

	// Declarar una variable del tipo de estructura esperado
	var response LocationResponse

	// Hacer la solicitud HTTP y llenar la estructura. Ahora con c.Previous, ya que salto el if, deberia estar con alguna direccion
	err := FetchData(c.Location.Previous, &response)
	if err != nil {
		return err
	}
	// Establecer los cambios en las rutas.
	c.Location.Previous = response.Previous
	c.Location.Next = response.Next

	// Acceder a los resultados
	for _, location := range response.Results {
		fmt.Println(location.Name)

	}

	return nil
}
