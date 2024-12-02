package main

import (
	"fmt"
)

//Atualização de Estrutura Aninhada: Crie uma struct Endereco com rua e numero, e uma struct
//Pessoa que contém nome, idade e endereco como campo. Escreva uma função que receba um
//ponteiro para Pessoa e altere o numero de Endereco

type Endereco struct {
	Rua    string
	Numero int
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func main() {
	slice := Pessoa{
		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Buteco",
			Numero: 18,
		},
	}

	novoEndereco(&slice)
}

func novoEndereco(slice *Pessoa) {
	sl := *slice

	sl.Endereco.Numero = 24

	fmt.Println(sl)
}
