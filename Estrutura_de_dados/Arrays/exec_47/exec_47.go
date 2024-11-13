package main

import (
	"fmt"
)

// leia um vetor de 10 números e os ordene em ordem crescente
func main() {
	numbers := []int{}
	var num int

	fmt.Println("Digite números inteiros ('0' para parar)")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}

		numbers = append(numbers, num)
	}

	// Bubble Sort
	for i := len(numbers); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println("Números em ordem crescente: ", numbers)

}
