package main

import (
	"fmt"
)

// altere o valor do segundo elemento de um slice. Imprima-o após a modificação.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice[1] = 7
	fmt.Println(slice)
}
