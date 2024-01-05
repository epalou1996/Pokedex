package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchData realiza una solicitud HTTP GET y decodifica la respuesta JSON en la estructura proporcionada
func FetchData(url *string, dataStructure interface{}) error {
	// Hacer la solicitud HTTP
	response, err := http.Get(*url)
	if err != nil {
		return fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	defer response.Body.Close()

	// Decodificar la respuesta JSON en la estructura proporcionada
	err = json.NewDecoder(response.Body).Decode(dataStructure)
	if err != nil {
		return fmt.Errorf("error al decodificar JSON: %v", err)
	}

	return nil
}
