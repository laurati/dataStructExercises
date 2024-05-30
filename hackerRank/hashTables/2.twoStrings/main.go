package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(twoStrings("hi", "world"))
}

func twoStrings(s1 string, s2 string) string {

	m := make(map[string]int)

	for _, v := range s1 {
		m[strconv.QuoteRune(v)]++
	}

	for _, v := range s2 {

		s := strconv.QuoteRune(v)

		if count, isThere := m[s]; count > 0 || isThere {
			return "Yes"
		}

	}
	return "No"
}
