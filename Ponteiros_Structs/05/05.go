package main

import "fmt"

//Alocação dinâmica de slices: crie uma função alocarSlice que aloca dinamicamente
//um slice de int com n elementos. Use essa função para criar slices de tamanhos
//diferentes e preencher com valores sequenciais (0,1,2,3,...)

func main() {
	sl1 := alocarSlice(5)
	sl2 := alocarSlice(7)
	sl3 := alocarSlice(13)

	fmt.Printf("\nSlice 1: %d", sl1)
	fmt.Printf("\nSlice 2: %d", sl2)
	fmt.Printf("\nSlice 3: %d", sl3)
}
func alocarSlice(n int) []int {
	sl := make([]int, n)
	//Alocação Dinâmica: A função make permite alocar o slice dinamicamente
	//em tempo de execução, com a capacidade de definir o tamanho (n) do slice

	//ALOCAÇÃO DINÂMICA: com a função make, o tamanho do slice é definido
	//apenas após a execução do programa, sendo alocado dinamicamente

	for i := 0; i < n; i++ {
		sl[i] = i
	}
	return sl
}
