package main

import "fmt"

// leia um vetor de 12 números inteiros e calcule a soma dos
// primos contidos nele
func main() {
	numbers := []int{}
	var num int
	somaPrimos := 0
	somaNaoPrimos := 0

	fmt.Println("Digite abaixo 12 números: ")
	for i := 0; i < 12; i++ {
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	// Itera pelo array e soma passando pela função de filtro
	for _, n := range numbers {
		if isPrime(n) {
			somaPrimos += n
		} else {
			somaNaoPrimos += n
		}
	}

	fmt.Printf("Soma dos números primos: %d\n", somaPrimos)
	fmt.Printf("Soma dos números não primos: %d\n", somaNaoPrimos)
}

// FUNCÃO DE FILTRO, retorna true se n é primo e false caso contrario. Os que retornarem
//true em caso, retornarão true na função e entrarão na soma, os demains serão não primos
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
