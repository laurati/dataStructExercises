package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, data chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	data := make(chan int)
	var wg sync.WaitGroup

	qtdWorkers := 10

	for i := 0; i < qtdWorkers; i++ {
		wg.Add(1)
		go worker(i, data, &wg)
	}

	go func() {
		for i := 0; i < 100; i++ {
			data <- i
		}
		close(data)
	}()
	wg.Wait()
}
