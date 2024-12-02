package main

import (
	"fmt"
)

//Alterando Valores em um Slice de Ponteiros:Crie um slice de ponteiros para inteiros e
//inicialize cada elemento com um valor sequencial (1, 2, 3, ...).Depois, altere todos os
//valores para o dobro usando um loop

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	dobrarSlice(&slice)
}

func dobrarSlice(slice *[]int) {

	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
	fmt.Println(*slice)
}
