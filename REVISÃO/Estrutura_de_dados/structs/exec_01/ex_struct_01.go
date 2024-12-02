package main

import (
	"fmt"
)

// Crie uma struct chamada Pessoa com os campos nome (string),
// idade (int) e altura (float64). Em seguida, inicialize uma variável
// do tipo Pessoa e imprima os valores de cada campo.
type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

// exec_04
// Adicione um campo cidade (string) à struct Pessoa. Atualize o slice do
// exercício_2 para incluir a cidade de cada pessoa. Em seguida, imprima o
// nome e a cidade de cada pessoa.
type NovaPessoa struct {
	Nome   string
	Idade  int
	Altura float64
	Cidade string
}
type Endereco struct {
	Rua    string
	Numero int
	cidade string
}
type Pessoa_com_Endereco struct {
	Nome     string
	Idade    int
	Altura   float64
	Cidade   string
	Endereco Endereco
}

func main() {
	idade_minima := 18
	pessoa := Pessoa{
		Nome:   "Carlos",
		Idade:  30,
		Altura: 1.80,
	}

	fmt.Println(pessoa)
	fmt.Println("Nome: ", pessoa.Nome)
	fmt.Println("Idade: ", pessoa.Idade)
	fmt.Println("Altura: ", pessoa.Altura)
	fmt.Println("\n-----------------\n")

	// exec_02
	//Usando a struct Pessoa do exercício anterior, crie um slice de Pessoa
	//e adicione três pessoas com valores diferentes. Em seguida, itere sobre
	//o slice e imprima o nome e idade de cada pessoa.

	slpessoa := []Pessoa{
		Pessoa{
			Nome:   "Maria",
			Idade:  17,
			Altura: 1.60,
		},
		Pessoa{
			Nome:   "Pedro",
			Idade:  25,
			Altura: 1.90,
		},
		Pessoa{
			Nome:   "Ana",
			Idade:  18,
			Altura: 1.70,
		}}

	for i, v := range slpessoa {
		fmt.Println(i, v)
	}
	media_idade(slpessoa)

	//exec_05
	//Crie uma nova struct chamada Endereco com os campos rua (string), numero (int) e
	//cidade (string). Modifique a struct Pessoa para incluir um campo endereco do tipo
	//Endereco. Crie uma pessoa com valores para o endereço eimprima os dados completos dela.
	pessoa_endereco := Pessoa_com_Endereco{
		Nome:   "Carlos",
		Idade:  30,
		Altura: 1.80,
		Endereco: Endereco{
			Rua:    "Avenida Paulista",
			Numero: 1000,
			cidade: "São Paulo",
		},
	}
	fmt.Println("\n-------------Struct Pessoa com endereço-------------")
	fmt.Println(pessoa_endereco)
	fmt.Println("Endereço: ", pessoa_endereco.Endereco)
	fmt.Println("Idade: ", pessoa_endereco.Idade)
	fmt.Println("Altura: ", pessoa_endereco.Altura)
	fmt.Println("Cidade: ", pessoa_endereco.Endereco.cidade)
	fmt.Println("Rua: ", pessoa_endereco.Endereco.Rua)
	fmt.Println("Numero: ", pessoa_endereco.Endereco.Numero)

	//exec_06
	//Crie um map onde a chave é o nome da pessoa (string) e o valor é do tipo Pessoa. Adicione
	//pelo menos três pessoas ao map e depois itere sobre ele, imprimindo o nome e a idade de cada pessoa.
	nome_pessoas := make(map[string]Pessoa)
	nome_pessoas["Maria"] = Pessoa{
		Nome:   "Maria",
		Idade:  20,
		Altura: 1.60,
	}
	nome_pessoas["Pedro"] = Pessoa{
		Nome:   "Pedro",
		Idade:  25,
		Altura: 1.90,
	}
	nome_pessoas["Ana"] = Pessoa{
		Nome:   "Ana",
		Idade:  22,
		Altura: 1.70,
	}
	fmt.Println("\n-------------Map Pessoa-------------")
	for i, v := range nome_pessoas {
		fmt.Println(i, v)
	}

	fmt.Println("\n-------------Maior idade-------------")
	printar_maior_idade := maior_idade(slpessoa, idade_minima)
	for i, v := range printar_maior_idade {
		fmt.Println(i, v)
	}

	// Chame o método Apresentacao e imprima o resultado.
	fmt.Println("------Metodo Apresentacao------", pessoa.Apresentacao())

	slPessoa_endereco := []Pessoa_com_Endereco{
		Pessoa_com_Endereco{
			Nome:   "Maria",
			Idade:  20,
			Altura: 1.60,
			Endereco: Endereco{
				Rua:    "Avenida Paulista",
				Numero: 1000,
				cidade: "São Paulo",
			},
		},
		Pessoa_com_Endereco{
			Nome:   "Pedro",
			Idade:  25,
			Altura: 1.90,
			Endereco: Endereco{
				Rua:    "Avenida Paulista",
				Numero: 1000,
				cidade: "São Paulo",
			},
		},
		Pessoa_com_Endereco{
			Nome:   "Ana",
			Idade:  22,
			Altura: 1.70,
			Endereco: Endereco{
				Rua:    "Avenida Paulista",
				Numero: 1000,
				cidade: "Manaus",
			}}}

	habitantes(slPessoa_endereco)

	fmt.Println("\n------Contagem de habitantes------\n", habitantes(slPessoa_endereco))
}

// exec_03
// Crie uma função que receba um slice de Pessoa e retorne a média de idade das
// pessoas. Use o slice criado no exercício 2 como entrada e imprima o resultado.
func media_idade(slpessoa []Pessoa) {
	var media int

	for _, v := range slpessoa {
		media += v.Idade
	}
	media /= len(slpessoa)
	fmt.Println("\nMedia de idade: ", media)
}

// exec_07
// Crie uma função que receba o slice de exec_02 e uma idade mínima. A função deve retornar
// um novo slice contendo apenas as pessoas que têm idade maior ou igual à idade mínima.
func maior_idade(slpessoa []Pessoa, idade_minima int) []Pessoa {
	pessoas_maior_idade := make([]Pessoa, 0)
	for _, v := range slpessoa {
		if v.Idade >= idade_minima {
			pessoas_maior_idade = append(pessoas_maior_idade, v)
		}
	}
	return pessoas_maior_idade
}

// exec_08
// Adicione um método Apresentacao à struct Pessoa que retorne uma string com o nome e a idade
// formatados, por exemplo: "Olá, meu nome é João e eu tenho 30 anos.". Crie uma pessoa, chame
// o método e imprima o resultado.
func (p Pessoa) Apresentacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

// exec_09
// Crie uma função que receba um slice de Pessoa e retorne um map onde as chaves são os nomes
// das cidades e os valores são a contagem de pessoas que vivem em cada cidade. Teste a função
// com o slice do exercício 2, incluindo as cidades no struct de Pessoa.
//
//	sliceLetras := make([]string, 2)
//	mapNomes := make(map[string]int)
func habitantes(slPessoa_endereco []Pessoa_com_Endereco) map[string]int {
	mapContagem_Cidades := make(map[string]int)

	for _, v := range slPessoa_endereco {
		mapContagem_Cidades[v.Endereco.cidade]++
	}
	return mapContagem_Cidades
}
