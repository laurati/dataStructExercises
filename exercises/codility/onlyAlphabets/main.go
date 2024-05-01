package main

import (
	"fmt"
	"strings"
)

// Given a String input which can have only alphabets.
// The string should be transformed if there is a combination of AB, CD, BA, DC
// by removing these occurrences and return the resultant String.

// example:

// Input : ABDCC -> Output: C
// Input : CABABD -> CABD -> CD -> Output: empty string

func Solution(S string) string {

	next := strings.ReplaceAll(S, "BA", "")

	for i := 0; i < len(S); i++ {
		next = strings.ReplaceAll(next, "BA", "")
		next = strings.ReplaceAll(next, "CD", "")
		next = strings.ReplaceAll(next, "AB", "")
		next = strings.ReplaceAll(next, "DC", "")
	}

	return next
}

func main() {
	S := "CABABD"
	fmt.Println(Solution(S))

}
