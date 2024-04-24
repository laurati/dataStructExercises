package main

import "fmt"

func main() {

	ch := make(chan string)

	go recebe("laura", ch)
	ler(ch)

}

// apenas enche o canal, canal recebe informações (receive-only)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// apenas esvazia o canal (send-only)
func ler(data <-chan string) {
	fmt.Println(<-data)
}
