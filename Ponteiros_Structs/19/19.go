package main

import (
	"fmt"
)

//Remoção de Elementos com Ponteiros: Implemente uma função removerElemento que recebe um
//ponteiro para um slice de int e um valor a ser removido, alterando o slice para remover
//todas as ocorrências do valor

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	removerElemento(&slice, 5)
}

func removerElemento(slice *[]int, valor int) {
	sl := *slice

	for i, v := range sl {
		if v == valor {
			sl = append(sl[:i], sl[i+1:]...)
		}
	}
	fmt.Println(sl)
}
