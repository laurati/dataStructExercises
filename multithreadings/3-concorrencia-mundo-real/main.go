package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// // nesse momento qnd acessar para alterar a var number, a mesma estará bloqueada
		// number++
		// // qnd terminar de atualizar o valor, é feito o unlock, e aí outras pessoas podem mudar o valor da var number
		// m.Unlock()
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})

	http.ListenAndServe(":3000", nil)

}
