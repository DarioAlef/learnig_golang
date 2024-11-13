package main

import (
	"fmt"
)

// Crie um array de inteiros com 5 elementos e altere o valor do segundo elemento.
func main() {
	array := []int{1, 2, 3, 4, 5}

	array[1] = 999

	fmt.Println(array)
}
