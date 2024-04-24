package main

import "time"

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
		time.Sleep(time.Second * 4)

	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	// esse for roda 2 vezes: na primeira printa o canal 1 e esvazia o canal, entao ele nao eh considerado na segunda iteração do for
	// na segunda printa o canal 2, pois ele ganha do do timeout
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1)

		case msg2 := <-c2:
			println("received", msg2)

		case <-time.After(time.Second * 3):
			println("timeout")

		}
	}
}
