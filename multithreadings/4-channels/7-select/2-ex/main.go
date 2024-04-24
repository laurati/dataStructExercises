package main

import "time"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	// colocando o canal antes do time, as vezes ele é printado primeiro do que o default
	go func() {
		c1 <- 1
		time.Sleep(time.Second * 4)

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

	default:
		println("default")
	}
}
