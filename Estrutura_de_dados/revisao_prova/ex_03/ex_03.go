package main

import (
	"fmt"
)

// Função que verifica se um número específico está presente em um array
func main() {
	array := []int{1, 2, 3, 4, 5}

	for i := range len(array) {
		if i == 3 {
			fmt.Println("O número", i, "está presente no array")
		}
	}
}
