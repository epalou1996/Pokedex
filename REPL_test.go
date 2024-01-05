package main

import "testing"

func TestCleanInputs(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "",
			output: []string{},
		},
		{
			input:  "hOla",
			output: []string{"hola"},
		},
		{
			input:  "Aduis mundo adsa 123",
			output: []string{"Aduis", "mundo", "adsa", "123"},
		},
	}
	for _, c := range cases {
		actual := cleanInputs(c.input)
		if len(actual) != len(c.output) {
			t.Errorf("El tamano de %v y %v no es el mismo", actual, c.output)
			continue
		}
	}

}
