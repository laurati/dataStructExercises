package main

func main() {

	// Channel: The channel's buffer is initialized with the specified buffer capacity.
	// If zero, or the size is omitted, the channel is unbuffered.
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
