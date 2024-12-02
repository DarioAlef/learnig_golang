package main

import (
	"fmt"
)

// escreva uma função recursiva que calcule o fatorial de um número. Além disso
// implemente uma versão iterativa e compare os tempos de execução.
func main() {
	var n int

	fmt.Println("\nDigite o número: ")
	fmt.Scan(&n)

	fmt.Printf("fatorial por recursão: %d", fatorial(n))
	fmt.Printf("\nfatorial por iteração: %d\n\n", fatItera(n))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fatorial(n-1)
	}
}

func fatItera(n int) int {
	fat_01 := 1
	for i := 1; i <= n; i++ {
		fat_01 *= i
	}
	return fat_01
}
