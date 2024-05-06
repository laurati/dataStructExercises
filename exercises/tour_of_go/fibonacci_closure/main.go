package main

import "fmt"

func main() {
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciRecursive(i))
	}
}

// 1)
// fibonacci is a function that returns
// a function that returns an int.
// func fibonacci() func() int {
// 	sum := 1
// 	sumAnterior := 0
// 	return func() int {
// 		// a := sum + sumAnterior
// 		// sumAnterior = sum
// 		// sum = a
// 		// return a

// 		// jeito mais inteligente
// 		a := sumAnterior

// 		sumAnterior, sum = sum, sumAnterior+sum

// 		return a
// 	}
// }

// 2)
// sem closure
// func fibonacci() {
// 	sumAnterior := 0
// 	sum := 1
// 	var a int
// 	for i := 0; i < 10; i++ {
// 		a = sumAnterior
// 		sumAnterior, sum = sum, sumAnterior+sum

// 		fmt.Println(a)
// 	}

// }

// func main() {

// 	fibonacci()

// }

// 3)
// recursiva
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
