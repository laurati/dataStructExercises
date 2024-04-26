package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		ch <- 4
	}()

	fmt.Println(<-ch)

}
