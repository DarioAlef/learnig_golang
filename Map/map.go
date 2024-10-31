package main

import (
	"fmt"
)

func main() {

	// MAP == DICT (python)

	mapNomes := make(map[string]int)
	fmt.Println("Map Nomes:", mapNomes)

	//Inserindo valores no mapNomes
	mapNomes["Raimundo"] = 37
	mapNomes["Maria"] = 18
	mapNomes["Ana"] = 20
	mapNomes["Paula"] = 22
	fmt.Println("Map Nomes:", mapNomes)

	//Criando um map literal sem 'make'
	mapProduto := map[int]string{
		1: "Resistor 1k Ohm 5%",
		2: "Diodo 1N4007",
		3: "LED difuso 5mm azul",
	}
	fmt.Println("\nMap Produto:", mapProduto)

	//Adicionando itens no mapProduto
	mapProduto[4] = "Capacitor"
	mapProduto[5] = "Placa montada"
	fmt.Println("Map Produto:", mapProduto)

	valor, existe := mapProduto[6]
	fmt.Println("\nMap Produto: ", valor, " - ", existe)

	//Removendo elementos do map com a função delete()
	delete(mapProduto, 4)
	fmt.Println("\nMap Produto:", mapProduto)

	//Iterando sobre o map usando 'range'
	for chave, valor := range mapProduto {
		fmt.Printf("\nChave: %d, Valor: %s\n", chave, valor)
	}

	//Tamanho de um map usando a função 'len()'
	fmt.Println("\nMap Produtos: ", len(mapProduto))
}
