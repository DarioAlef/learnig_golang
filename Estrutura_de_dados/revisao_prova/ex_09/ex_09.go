package main

import "fmt"

//Crie dois slices de inteiros, um com 3 elementos e outro com 2 elementos.
//Concatenate os dois slices e imprima o resultado.
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}
