package main

import (
	"fmt"
)

func main() {
	arrayNomes := [4]string{"João", "Maria", "Pedro", "Ana"}
	fmt.Println(arrayNomes)

	fmt.Println("Tamanho do array é: ", len(arrayNomes))

	arrayNumeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		fmt.Println(arrayNumeros[i])
	}
	fmt.Printf("\n")

	//   MATRIZES

	array2D := [4][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D[i]); j++ {
			fmt.Println(array2D[i][j])
		}
		fmt.Println("\n")
	}
}
