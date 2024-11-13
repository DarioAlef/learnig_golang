package main

import (
	"fmt"
)

// preencha um vetor com os 10 primeiros números primos
func main() {
	numbers := []int{}

	for i := 2; len(numbers) < 10; i++ {
		if primos(i) {
			numbers = append(numbers, i)
		}
	}
	fmt.Println("\nOs dez primeiros números primos ->")
}

func primos(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; n >= i*i; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	numbers := []int{}

// 	for i := 2; len(numbers) < 10; i++ {
// 		if isPrime(i) {
// 			numbers = append(numbers, i)
// 		}
// 	}

// 	fmt.Println("Primeiros 10 números primos: ", numbers)
// }

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	}
// 	if n == 2 || n == 3 {
// 		return true
// 	}
// 	if n%2 == 0 || n%3 == 0 {
// 		return false
// 	}
// 	for i := 5; i*i <= n; i += 6 {
// 		if n%i == 0 || n%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
