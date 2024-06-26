package main

import "fmt"

// duas funcs enviando x valores cada uma pra um canal
// for loop x valores, select case <-x

func main() {

	a := make(chan int)
	b := make(chan int)

	x := 500

	go func(y int) {
		for i := 0; i < y; i++ {
			a <- i
		}
	}(x)

	go func(y int) {
		for i := 0; i < y; i++ {
			b <- i
		}
	}(x)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("canal A: ", v)
		case v := <-b:
			fmt.Println("canal B: ", v)
		}
	}
}
