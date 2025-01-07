package main

import "fmt"

func main() {

	str := "Лорем ипсум долор сит амет, еум алияуип яуаерендум ут, нам дицта орнатус ут"

	var result1 = make(chan map[string]int)
	var result2 = make(chan string)
	go countStr(str, result1)
	go invertStr(str, result2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-result1:
			fmt.Println(msg)
		case msg := <-result2:
			fmt.Println(msg)

		}
	}

}

func countStr(str string, ch chan map[string]int) {
	var ocurrencies = make(map[string]int, 0)
	for _, s := range str {
		ocurrencies[string(s)] += 1
	}

	ch <- ocurrencies
}

func invertStr(str string, ch chan string) {
	var result string
	newstr := []rune(str)
	for i := len(newstr) - 1; i > 0; i-- {
		result += string(newstr[i])
	}

	ch <- result
}
