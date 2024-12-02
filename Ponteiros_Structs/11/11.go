package main

import (
	"fmt"
)

//Slice de Structs e Ponteiros: Crie uma struct Produto com campos nome e preco. Faça uma função
//que receba um ponteiro, apontando para um slice de Produto e altere os preços para um desconto de 10%

type Produto struct {
	Nome  string
	Preco float32
}

func main() {
	slP := Produto{
		Nome:  "Amaciante",
		Preco: 16.5,
	}

	desconto(&slP)
}

func desconto(sl *Produto) {
	sl.Preco = sl.Preco * 0.9

	fmt.Println(sl)
}
