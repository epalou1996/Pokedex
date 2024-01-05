package main

import (
	"fmt"
	"os"
)

// Mata el REPL
func commandExit(c *config) error {
	fmt.Println("Hasta luego!")
	os.Exit(0)
	return nil
}
