package main

import "fmt"

//Crie uma struct Endereço com os campos rua, numero, cidade. Crie uma struct Pessoa
//que tenha o campo endereço. Instancia e imprima o endereço completo

type Endereço struct {
	Rua    string
	Numero int
	Cidade string
}
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereço
}

func main() {
	sl := Pessoa{
		Nome:  "Joaquim",
		Idade: 20,
		Endereco: Endereço{
			Rua:    "Av. Paulista",
			Numero: 10,
			Cidade: "Sao Paulo",
		}}

	fmt.Printf("\n%+v\n\n", sl)
}
