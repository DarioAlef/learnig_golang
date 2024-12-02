package main

import "fmt"

//Implemente recursão para calcular o fatorial de um número n. Use um map
//para armazenar os valores já calculados, evitando repetições de cálculos.

func main() {
	var n int
	map_fatorial := make(map[int]int)

	fmt.Println("Digite um número: ")
	fmt.Scan(&n)
	map_fatorial[n] = fatorial(n)
	fmt.Println(map_fatorial)
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fatorial(n-1)
	}
}
