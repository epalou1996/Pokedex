package main

func main() {
	// Creacion de configuraciones para el app e iniciacion.
	next := "https://pokeapi.co/api/v2/location-area"
	previous := ""
	cfg := &config{
		Location: MapsURLs{
			Next:     &next,
			Previous: &previous,
		},
	}
	StartRepl(cfg)
}
