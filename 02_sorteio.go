package main

import (
	"fmt"
	"math/rand"
)

//' * ' é usado para acessar o valor de referência de um ponteiro

func main() {
	n := rand.Intn(101) //rand.Intn gera um valor aleatório entre 0 e 100

	var tentativa int

	fmt.Println("Tente advinhar o número entre 0 e 100: ")
	fmt.Scan(&tentativa)

	for tentativa != n {
		if tentativa < n {
			fmt.Println("Tente um número maior")
		} else {
			fmt.Println("Tente um número menor")
		}
		fmt.Scan(&tentativa)
	}
	fmt.Println("Parabéns, você acertou!")

	// for i := 20; i < 30; i++ {   //for usado como for
	// 	fmt.Println(i)
	// }

	// I := 0
	// for I < 10 {   //for usado como while
	// 	fmt.Println(I)
	// 	I++
	// }
}

// func sortear(n int) {}
