package main

import "fmt"

//Desenvolva uma função recursiva que recebe uma string e retorna
//a mesma string invertida. Não utilize loops, apenas recursão.
func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	inverter(palavra)
	fmt.Println(inverter(palavra))
}
func inverter(palavra string) string {

	if len(palavra) <= 1 {
		return palavra
	}
	// Recursivamente, 'palavra[1:]' remove o primeiro caractere
	//e 'string(palavra[0])' coloca o primeiro caractere no final
	return inverter(palavra[1:]) + string(palavra[0])

}

// func fatorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	} else {
// 		return n * fatorial(n-1)
// 	}
// }

//for i := len(palavra) - 1; i >= 0; i-- {
//	reverso += string(palavra[i])

//i começa pelo tamanho de palavra (-1) para ser ao contrário
//adicionando à variável array 'reverso' }
