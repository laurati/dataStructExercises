package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	contador := 0
	totalDeGoroutines := 10

	wg.Add(totalDeGoroutines)

	for i := 0; i < totalDeGoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			// funciona como o sleep nbesse caso
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
