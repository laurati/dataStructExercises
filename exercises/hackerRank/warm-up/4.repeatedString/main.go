package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "aba"
	n := 10

	searchChar := "a"

	sLength := len(s)                 // tamanho da string
	repetitionComplete := n / sLength // qts vezes a string completa cabe no n
	remainder := n % sLength          // restante dos caracteres apos repeticoes completas

	aCount := strings.Count(s, searchChar)*int(repetitionComplete) + strings.Count(s[:remainder], searchChar)
	fmt.Println(aCount)
}
