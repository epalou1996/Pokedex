package main

func main() {
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
