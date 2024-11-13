package main

import "fmt"

//função que remova o 2° elemento de um slice e retorne o novo slice.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	remove(slice, 1)
	fmt.Println(slice)
}

func remove(slice []int, index int) []int {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
