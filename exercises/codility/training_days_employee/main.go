package main

import (
	"fmt"
	"strings"
)

func Solution(E []string) int {
	maxAttendees := 0

	// Iterar sobre todos os pares de dias possíveis
	for day1 := 0; day1 <= 9; day1++ {
		for day2 := day1 + 1; day2 <= 9; day2++ {
			// Contar quantos funcionários podem comparecer em pelo menos um dos dois dias
			attendees := 0
			for _, preferences := range E {
				if strings.ContainsRune(preferences, rune(day1+'0')) || strings.ContainsRune(preferences, rune(day2+'0')) {
					attendees++
				}
			}
			// Atualizar o número máximo de funcionários que podem comparecer
			if attendees > maxAttendees {
				maxAttendees = attendees
			}
		}
	}

	return maxAttendees
}

func main() {
	E := []string{"039", "4", "14", "32", "", "34", "7"}
	// E := []string{"5421", "245", "1452", "0345", "53", "354"}
	fmt.Println(Solution(E))
}
