package main

import "fmt"

//Escreva uma função recursiva que conte quantos elementos
//são maiores que 10 em um slice de inteiros.
func main() {
	slice := make([]int, 0)
	var n int

	fmt.Println("Digite alguns números ('0' para parar): ")

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		slice = append(slice, n)
	}
	fmt.Println("Quantidade de números maiores que 10: ", contag(slice))
}

func contag(slice []int) []int {
	maiores := []int{}

	for i:= range slice {
		if slice[i] > 10 {
			maiores = append(maiores, slice[i])
		}
	}
	return maiores
}