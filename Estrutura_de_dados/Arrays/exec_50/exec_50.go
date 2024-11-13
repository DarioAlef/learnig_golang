package main

import "fmt"

// copie todos os elementos de um vetor para outro sem copiar
// os números repetidos
func main() {
	numbers := []int{}
	var num int

	fmt.Println("Digite um número (0 para parar): ")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}

	//slice começando em 0 e indo até o tamanho de 'numbers'
	uniqueNumbers := make([]int, 0, len(numbers))
	uniqueMap := make(map[int]bool)

	for _, num := range numbers {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	fmt.Println("Números unicos: ", uniqueNumbers)
}
