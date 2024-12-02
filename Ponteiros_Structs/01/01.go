package main

import "fmt"

//Alteração por referência: escreva uma função 'alterarValor' que recebe um
//ponteiro para um int e altera o valor para o dobro

func main() {
	x := 10

	p := &x
	fmt.Println(alterarValor(p))
}

func alterarValor(p *int) int {
	n := 2 * (*p)

	return n
}
