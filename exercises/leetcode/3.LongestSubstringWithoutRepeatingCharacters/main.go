package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	// Map para armazenar o último índice visto de cada caractere
	lastSeen := make(map[byte]int)

	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {

		idx, found := lastSeen[s[i]]

		if found && idx >= start {
			start = idx + 1
		}

		lastSeen[s[i]] = i

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

	}
	return maxLength

}

func main() {
	s := "abcabcbb"
	maxLength := lengthOfLongestSubstring(s)

	fmt.Println(maxLength)
}
