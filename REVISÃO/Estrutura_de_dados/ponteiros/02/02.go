package main

import "fmt"

//Crie uma função que inverta os elementos de um slice de inteiros usando ponteiros
//para modificar o slice original.

func main() {

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(slice)
	// for i := 0; i < len(slice)/2; i++ {
	// 	slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	// }
	// fmt.Println(slice)

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sl)

	var psl *[]int = &sl
	//Declara a variável psl como um ponteiro para um slice 'sl' e atribui a
	//ela o endereço de memória de 'sl' usando o operador '&'. Agora, psl
	//"aponta" para o slice sl, qualquer modificação feita em *psl refletirá em sl.
	fmt.Println(invert(psl))
}

func invert(psl *[]int) []int {
	vpsl := *psl

	for i := 0; i < len(vpsl)/2; i++ {
		(vpsl)[i], (vpsl)[len(vpsl)-i-1] = (vpsl)[len(vpsl)-i-1], (vpsl)[i]
	}
	return *psl
}
