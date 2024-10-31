package main

import (
	"fmt"
	"strconv"
)

var aux string

func main() {
	//   MATRIZES

	cartela()
	escolhidos()
	cartela()
}

func escolhidos() {
	var aux int
	//numeros_escolhidos := [6]int

	for i := 0; i < 6; i++ {
		fmt.Printf("\nEscolha um número de 1 a 60: ")
		fmt.Scanln(&aux)
	}
	return
}

func cartela() {
	arrayLoto := [6][10]string{}
	count := 1

	for i := 0; i < len(arrayLoto); i++ {
		for j := 0; j < len(arrayLoto[i]); j++ {
			arrayLoto[i][j] = strconv.Itoa(count)

			if count < 10 {
				if aux == arrayLoto[i][j] {
					fmt.Printf(" [0%v] ", arrayLoto[i][j])

				} else {
					fmt.Printf(" 0%v ", arrayLoto[i][j])
				}
				count++
			} else {
				if aux == arrayLoto[i][j] {
					fmt.Printf(" [%v] ", arrayLoto[i][j])

				} else {
					fmt.Printf(" %v ", arrayLoto[i][j])
				}
				count++

			}
		}
		fmt.Print("\n")
	}
	return
}
