package main

import (
	"fmt"
)

// Crie um slice de 6 inteiros. Extraia um novo slice contendo os três
// últimos elementos do slice original.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	novo_slice := slice[3:]
	fmt.Println(novo_slice)
}
