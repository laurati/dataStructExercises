package main

import "fmt"

func main() {

	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	checkMagazine(magazine, note)

	magazine2 := []string{"two", "times", "three", "is", "not", "four"}
	note2 := []string{"two", "times", "two", "is", "four"}
	checkMagazine(magazine2, note2)

}

func checkMagazine(magazine []string, note []string) {

	wordCount := make(map[string]int)

	for _, v := range magazine {
		wordCount[v]++
	}

	for _, word := range note {

		count, exists := wordCount[word]
		if count == 0 || !exists {
			fmt.Println("No")
			return
		}
		wordCount[word]--
	}
	fmt.Println("Yes")
}
