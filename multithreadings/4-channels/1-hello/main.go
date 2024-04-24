package main

import "fmt"

// Thread 1
func main() {

	canal := make(chan string) //vazio

	// Thread 2
	go func() {
		// pego a string e preencho o canal
		canal <- "ola mundo" //cheio
	}()

	// Thread 1
	// o valor que add no canal, vou jogar pra var msg
	msg := <-canal //canal esvazia
	fmt.Println(msg)
}
