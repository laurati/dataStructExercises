package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Ex: duas goroutines
// func main() {

// 	wg.Add(2)

// 	go func() { fmt.Println("Goroutine número 1"); wg.Done() }()
// 	go func() { fmt.Println("Goroutine número 2"); wg.Done() }()

// 	wg.Wait()
// }

// Ex: nº variado de goroutines
func main() {
	variasGoroutines(10)
	wg.Wait()
}

func variasGoroutines(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func(i int) {
			fmt.Println("Goroutine número: ", i+1)
			wg.Done()
		}(i)
	}
}
