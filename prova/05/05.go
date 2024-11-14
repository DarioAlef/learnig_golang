package main

import "fmt"

//Crie uma função que calcule a soma das diagonais principais e secundárias
//de uma matriz 3x3 e retorne o resultado
func main() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}
	diagonais(matriz)
}

func diagonais(matriz [3][3]int) {
	var dp int
	var ds int
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			if i == j {
				dp += matriz[i][j]
			}
			if i+j == 2 {
				ds += matriz[i][j]
			}
		}
	}
	fmt.Printf("\nDiagonal principal: %d", dp)
	fmt.Printf("\nDiagonal secundária: %d\n\n", ds)
}
