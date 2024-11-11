package main

import "fmt"

func main() {
	matriz1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	transposta(matriz1)
	fmt.Println(transposta(matriz1))
}
func transposta(matriz1 [3][3]int) [3][3]int {

	//Iterar sobre a matriz
	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz1); j++ {
			matriz1[i][j] = matriz1[j][i]
		}
	}
	return matriz1
}
