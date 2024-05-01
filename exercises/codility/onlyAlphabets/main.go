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
	next1 := strings.ReplaceAll(next, "CD", "")
	next2 := strings.ReplaceAll(next1, "AB", "")
	next3 := strings.ReplaceAll(next2, "DC", "")
	return next3
}

func removePairs(input string) string {
	result := make([]byte, 0, len(input))

	for i := 0; i < len(input); i++ {
		if i < len(input)-1 &&
			(input[i] == 'A' && input[i+1] == 'B' || input[i] == 'C' && input[i+1] == 'D' ||
				input[i] == 'B' && input[i+1] == 'A' || input[i] == 'D' && input[i+1] == 'C') {
			i++ // Skip next character
		} else {
			result = append(result, input[i])
		}
	}

	return string(result)
}

func main() {
	S := "ABDCC"
	fmt.Println(Solution(S))

	fmt.Println(removePairs("ABDCC"))
}
