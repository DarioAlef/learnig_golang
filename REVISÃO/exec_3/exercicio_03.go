package main

import "fmt"

// soliocite uma lista do usuário e faça a soma dos números pares
func main() {
	fmt.Println("Digite os números (negativo para parar:")
	var numeros []int
	var pares int

	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		numeros = append(numeros, num)
	}
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			pares = pares + numeros[i]
		}
	}
	fmt.Println("\nA soma de todos os números pares é: ", pares)
}
