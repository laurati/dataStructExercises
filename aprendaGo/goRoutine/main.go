package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// numero de processadores
	fmt.Println(runtime.NumCPU())

	// numero de goroutines
	// começo o programa com uma goroutine que é a main
	fmt.Println(runtime.NumGoroutine())

	// o programa vai esperar 2 goroutine
	wg.Add(2)

	// as duas funcs rodam de maneira concorrente e ao mesmo tempo
	go print1(&wg)
	go print2()

	// numero de goroutines
	// programa roda as 3 goroutines
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func print1(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("print1: %d\n", i)
		time.Sleep(10)
	}
	wg.Done()
}

func print2() {
	for i := 0; i < 50; i += 10 {
		fmt.Printf("print2: %d\n", i)
		time.Sleep(10)
	}
	wg.Done()
}
