package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	x := strings.Fields(s)
	m := map[string]int{}
	for _, v := range x {
		m[v]++
	}

	return m

}

func main() {

	a := WordCount("I am Laura and I hate donut")
	fmt.Println(a)
}
