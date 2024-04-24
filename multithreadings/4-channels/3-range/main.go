package main

import "fmt"

// Thread 1
func main() {

	ch := make(chan int)
	go pusblish(ch)
	reader(ch)

}

func reader(ch chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}

func pusblish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
