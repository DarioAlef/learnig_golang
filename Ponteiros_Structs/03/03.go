package main

import (
	"fmt"
)

//Structs e Ponteiro: crie uma struct PEssoa com campos 'nome' e 'idade'. Escreva uma função
//que receba um ponteiro para Pessoa e altere a idade para um valor passado como argumento

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	age := 20

	sl := Pessoa{
		Nome:  "João",
		Idade: 30,
	}

	fmt.Println("\nOriginal:", sl)
	swapAge(&sl, age)
	fmt.Print("Alterado: ", sl, "\n\n")

}

func swapAge(p *Pessoa, age int) {
	p.Idade = age

}
