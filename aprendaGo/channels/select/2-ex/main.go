package main

import "fmt"

// func1 recebe x valores de canal, depois manda qualquer coisa pra chan quit
// func2 for infinito, select: case envial pra canal, case recebe de quit

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 5; i++ {
		// Aguarda e imprime a mensagem recebida do canal
		fmt.Println("Recebido", <-canal)
	}
	// Após receber 5 mensagens, envia um sinal de conclusão para o canal 'quit'.
	quit <- 0
}

func enviaPraCanal(canal chan int, quit chan int) {
	a := 1
	for {
		select {
		// Se o canal 'canal' estiver pronto para receber, envia o valor de 'a' para ele e incrementa 'a'.
		case canal <- a:
			a++
		// Se um valor for recebido no canal 'quit', retorna e encerra a goroutine.
		case <-quit:
			return
		}
	}
}
