package main

import (
	"fmt"
)

// Crie um programa com o seguinte memnu:
// 1. Adicionar nome e sexo de uma pessoae
// 2. Listar todas os nomes
// 3. Total de homens e total de mulheres
// 4. Sair
func main() {
	var opcao int

	for {
		fmt.Println("\nMenu: (selecione a opção desejada)")
		fmt.Println("1. Adicionar nome e sexo de uma pessoa")
		fmt.Println("2. Listar todas os nomes")
		fmt.Println("3. Total de homens e total de mulheres")
		fmt.Println("4. Sair")
		fmt.Scan(&opcao)

		escolha(opcao)
	}
}

func escolha(opcao int) {
	var nome string
	var sexo string
	mapPessoas := make(map[string]string)

	switch opcao {
	case 1:
		for {
			fmt.Println("Digite o nome da pessoa: ")
			fmt.Scan(&nome)
			if nome == "n" {
				break
			}
			fmt.Println("Digite o sexo da pessoa: ")
			fmt.Scan(&sexo)
			if sexo == "n" {
				break
			}
		}
		for chave, valor := range mapPessoas {
			fmt.Println(chave, valor)
		}
	case 2:
		for chave, valor := range mapPessoas {
			fmt.Println(chave)
			fmt.Println(valor)
		}
	}
}
