package main

import "fmt"

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}

	// m := make(map[int]int)
	// sobra := 0
	// for _, v := range A {
	// 	m[v]++
	// 	if m[v] < 2 {
	// 		sobra = v
	// 	}

	// }

	// fmt.Println(m, sobra)

	// ==================================================
	// XOR

	// Encontrar o elemento sem par em um array
	unpaired := 0
	for _, v := range A {
		unpaired ^= v
	}
	fmt.Println(unpaired)

	// Trocar valores de variáveis sem usar uma variável temporária
	a := 5
	b := 10
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)

	// verificar se é par ou impar
	num := 6
	see := num&1 == 1 //Verifica se o último bit é 1,; impar -> true, par -> false
	fmt.Println(see)

	// Verificar se dois números são diferentes
	dif := a^b != 0 // Verifica se o resultado do XOR é diferente de zero; true (5 e 10 são diferentes)
	fmt.Println(dif)
}
