package main

import (
	"fmt"
)

// Crie um slice vazio e adicione os números de 1 a 5 usando o operador
// append. Imprima o slice resultante.
func main() {
	slice := []int{}

	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)
}
