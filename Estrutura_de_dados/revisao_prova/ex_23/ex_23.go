package main
import (
	"fmt"
)
//Escreva uma função que faça uma cópia de um slice de inteiros e retorne o 
//novo slice. Modifique o slice original e verifique se a cópia é afetada
func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//novoSlice := slCopia(sl)
	fmt.Println(slCopia(sl))

	// sl[0] = 0

	// for _, v := range novoSlice {
	// 	for _, v2 := range sl {
	// 		if v != v2 {
	// 			fmt.Println("Os slices não são iguais")
	// 		}
	// 	}
		
	// } 
}

func slCopia(sl []int) []int {
	novoSl := make([]int, len(sl))
	for i, v := range sl {
		novoSl[i] = v
	}
	return novoSl
}