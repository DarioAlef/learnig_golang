package main

import "fmt"

//Crie um array com 10 inteiros. Converta-o para um slice
//contendo os primeiros 5 elementos e imprima esse slice.
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := array[0:5]

	fmt.Println(slice)
}
