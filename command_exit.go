package main

import (
	"fmt"
	"os"
)

// Mata el REPL
func commandExit(c *config, args ...string) error {
	fmt.Println("Hasta luego!")
	os.Exit(0)
	return nil
}
