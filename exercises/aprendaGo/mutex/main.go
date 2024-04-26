package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Ex1: race condition
// dessa forma cada vez vai atualizar o contador com um numero diferente
// varias threads acessaram o contador ao mesmo tempo gerando race condition
// func main() {
// 	contador := 0
// 	totalDeGoroutines := 10

// 	var wg sync.WaitGroup
// 	wg.Add(totalDeGoroutines)

// 	for i := 0; i < totalDeGoroutines; i++ {
// 		go func() {
// 			v := contador
// 			// funciona como o sleep nbesse caso
// 			runtime.Gosched()
// 			v++
// 			contador = v
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(contador)
// }

// Ex2: mutex
func main() {

	fmt.Println("Goroutines antes", runtime.NumGoroutine())

	contador := 0
	totalDeGoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totalDeGoroutines)

	var mu sync.Mutex

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			time.Sleep(15)
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines dentro do for", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines final", runtime.NumGoroutine())
	fmt.Println(contador)
}
