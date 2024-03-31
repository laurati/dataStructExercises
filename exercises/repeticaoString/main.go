package main

import "fmt"

func countEachChar(s string) {
	mapStr := make(map[string]int)
	for i := 0; i < len(s); i++ {
		mapStr[string(s[i])]++
	}
	fmt.Println(mapStr)
}

func countEachWord(s []string) {
	mapStr := make(map[string]int)
	for _, v := range s {
		mapStr[v]++
	}
	fmt.Println(mapStr)
}

func main() {

	countEachChar("abcabcbb")

	s := []string{"ana", "camila", "laura", "ana", "bia", "bia", "laura"}
	countEachWord(s)
}
