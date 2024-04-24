package main

import "time"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 4)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		println("received", msg1)

	case msg2 := <-c2:
		println("received", msg2)

	case <-time.After(time.Second * 3):
		println("timeout")
	}
}
