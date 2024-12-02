package main

import "fmt"

//Crie uma struct Endereco com campos rua, numero e cidade. Crie uma struct Pessoa que tenha
//um campo Endereco. Inicialize uma instância da struct Pessoa e imprima o endereço completo

type Endereco struct {
	cidade, rua string
	numero      int
}
type Pessoa struct {
	Nome     string
	idade    int
	Endereco Endereco
}

func main() {
	sl := []Pessoa{
		{
			Nome:  "Maria",
			idade: 25,
			Endereco: Endereco{
				cidade: "Sao Paulo",
				rua:    "Av Copacabana",
				numero: 20,
			}}}

	fmt.Println(sl)
}
