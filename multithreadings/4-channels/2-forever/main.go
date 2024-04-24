package main

// Thread 1
func main() {

	forever := make(chan bool) //vazio

	// Thread 2
	go func() {
		for i := 0; i < 5; i++ {
			print(i)
		}
		forever <- true
	}()

	<-forever
}
