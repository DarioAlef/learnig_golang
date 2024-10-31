package main

import (
	"fmt"
)

func main() {
	//Slice: criando diretamente do Array
	sliceNumeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice Números: ", sliceNumeros)

	//Slice: criando vazio
	sliceLetras := make([]string, 2)
	fmt.Println("Slice Letras: ", sliceLetras)

	//Add elemento ao SliceNumeros

	sliceNumeros = append(sliceNumeros, 6)
	fmt.Println("Slice Numeros: ", sliceNumeros)

	//Add elementos ao sliceLetras
	sliceLetras = append(sliceLetras, "A")
	sliceLetras = append(sliceLetras, "B")
	sliceLetras = append(sliceLetras, "C")
	sliceLetras = append(sliceLetras, "D")
	sliceLetras = append(sliceLetras, "E")
	fmt.Println("Slice Letras: \n", sliceLetras)

	//Iterando slice com 'for'
	for i := 0; i < len(sliceNumeros); i++ {
		fmt.Println("Slices Numeros com 'for': ", sliceNumeros[i])
	}

	//Iterando slice com for e range
	fmt.Println("\nSlice Letras com 'range':")
	for indice, valor := range sliceLetras {
		fmt.Printf("Índice: %d, Valor: %s\n", indice, valor)
	}

	//Removendo elementos do sliceNumeros:
	fmt.Println("\nRemovendo elementos do sliceNumeros: ")
	sliceNumeros = append(sliceNumeros[:2], sliceNumeros[2:]...)
	//eu removi o elemento entre os índices 2 e 3, dividindo o slice e juntando novamente,
	//assim fazendo com que o espaço em branco fosse eliminado
	fmt.Println("Slice numeros: ", sliceNumeros)

	//Tamanho do Slice utilizando len
	fmt.Println("\nTamanho do Slice utilizando len: ")
	fmt.Printf("Tamanho do Slice é [%d] \n", len(sliceNumeros))
}
