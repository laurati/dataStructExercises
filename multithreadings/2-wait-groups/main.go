package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {

	waitGroup := sync.WaitGroup{}

	// cada vez que ele passa pelo wg.Done() no for, ele diminui uma unidade do add e
	// o waitGroup.Wait() segura a execução da main
	// qnd terminar as operações, ele saberá que não precisa mais esperar
	waitGroup.Add(25)

	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
