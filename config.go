package main

// Config holds the URLs for paging through location areas.
type Config struct {
	Next     string
	Previous string
}

// initialURL is where we start (first 20 areas).
const initialURL = "https://pokeapi.co/api/v2/location-area?limit=20"
