package main

import (
	"fmt"
	"iter"
)

// Iterador que divide uma string em partes
func Split(s, sep string) iter.Seq[string] {
	return func(yield func(string) bool) {
		start := 0
		for i := range s {
			if string(s[i]) == sep {
				if !yield(s[start:i]) {
					return
				}
				start = i + 1
			}
		}
		if start < len(s) {
			yield(s[start:])
		}
	}
}

func main() {
	// Divide a string e itera sobre as partes
	for part := range Split("oi amor lindo", " ") {
		fmt.Println(part)
	}
}
