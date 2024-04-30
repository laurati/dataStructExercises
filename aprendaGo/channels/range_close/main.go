package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	wg.Add(2)

	go meuloop(10, c)
	go printChannel(c)
	wg.Wait()
}

func meuloop(t int, ch chan<- int) {
	for i := 0; i < t; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func printChannel(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido do canal: ", v)
	}
	wg.Done()
}
