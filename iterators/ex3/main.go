package main

import (
	"fmt"
	"iter"
)

// Define tipos
type City struct {
	Name string
}

type Country struct {
	cities    []*City
	languages []string
}

// Iterador de cidades
func (c *Country) Cities() iter.Seq[*City] {
	return func(yield func(*City) bool) {
		for _, city := range c.cities {
			if !yield(city) {
				return
			}
		}
	}
}

// Iterador de idiomas
func (c *Country) Languages() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, lang := range c.languages {
			if !yield(lang) {
				return
			}
		}
	}
}

func main() {
	// Cria um país
	country := Country{
		cities: []*City{
			{Name: "New York"},
			{Name: "Los Angeles"},
			{Name: "Chicago"},
		},
		languages: []string{"English", "Spanish"},
	}

	// Itera sobre as cidades
	for city := range country.Cities() {
		fmt.Println("City:", city.Name)
	}

	// Itera sobre os idiomas
	for lang := range country.Languages() {
		fmt.Println("Language:", lang)
	}
}
