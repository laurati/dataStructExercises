package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(10)

	// Thread 2
	go pusblish(ch)

	// Thread 3
	go reader(ch, &wg)

	wg.Wait() // segura o processo da main
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Println(x)
		wg.Done()
	}
}

func pusblish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// close(ch)
}
