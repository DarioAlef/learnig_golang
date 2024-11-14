package main 

import ("fmt")	
//Crie uma função que compare duas matrizes 2x2 e retorne 
//true se forem iguais e false caso contrário.
func main () {
	matriz1 := [2][2]int{{1, 2}, {3, 4}}
	matriz2 := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(compare(matriz1, matriz2))
}

func compare(matriz1 [2][2]int, matriz2 [2][2]int) bool {
	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz1); j++ {
			if matriz1[i][j] == matriz2[i][j] {
				return true
			}
		}
	}
	return false
}

