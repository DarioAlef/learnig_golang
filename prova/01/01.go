package main

import (
	"fmt"
)

//ordene um array de 10 posições de forma invertida

func main() {
	array := []int{}
	var n int

	fmt.Println("\nDigite 10 números inteiros")
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)

		array = append(array, n)
	}
	//bubble sort invertido
	for i := len(array); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
