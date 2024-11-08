package main

import (
	"fmt"
)

type Pessoa struct {
	Nome     string
	Idade    int
	Altura   float64
	Endereco Endereco
}
type Endereco struct {
	Cidade string
	Rua    string
	Numero int
}

func main() {
	idade_minima := 18

	slpessoa := []Pessoa{
		Pessoa{
			Nome:   "Carlos",
			Idade:  30,
			Altura: 1.80,
			Endereco: Endereco{
				Cidade: "São Paulo",
				Rua:    "Av. Paulista",
				Numero: 10,
			},
		},
		Pessoa{
			Nome:   "Maria",
			Idade:  17,
			Altura: 1.60,
			Endereco: Endereco{
				Cidade: "São Paulo",
				Rua:    "Av. Copacabana",
				Numero: 20,
			},
		},
		Pessoa{
			Nome:   "Pedro",
			Idade:  25,
			Altura: 1.90,
			Endereco: Endereco{
				Cidade: "São Paulo",
				Rua:    "Rua da Praia",
				Numero: 30,
			},
		},
		Pessoa{
			Nome:   "João",
			Idade:  18,
			Altura: 1.70,
			Endereco: Endereco{
				Cidade: "São Paulo",
				Rua:    "Av. Europa",
				Numero: 40,
			},
		},
		Pessoa{
			Nome:   "Dário",
			Idade:  23,
			Altura: 1.80,
			Endereco: Endereco{
				Cidade: "Manaus",
				Rua:    "Domingo Russo",
				Numero: 160,
			}}}

	apresentacao(slpessoa)
	fmt.Println(apresentacao(slpessoa))

	media_idade(slpessoa)
	fmt.Printf("\nA média de idade é: %d", media_idade(slpessoa))

	maior_idade(slpessoa, idade_minima)
	fmt.Printf("\n\nAs pessoas maiores de idade são: %v", maior_idade(slpessoa, idade_minima))

	habitantes(slpessoa)
	fmt.Printf("\n\nA contagem de pessoas por cidade é: %v\n", habitantes(slpessoa))
}

func media_idade(slpessoa []Pessoa) int {
	var media_de_idade int

	//irá iterar no slice slpessoa, mas a atribuição será apenas com idade
	//usando '.Idade'
	for _, v := range slpessoa {
		media_de_idade += v.Idade
	}
	media_de_idade /= len(slpessoa)

	return media_de_idade
}

func maior_idade(slpessoa []Pessoa, idade_minima int) []Pessoa {
	maior_de_idade := make([]Pessoa, 0)

	for _, v := range slpessoa {
		if v.Idade >= idade_minima {
			maior_de_idade = append(maior_de_idade, v)
		}
	}
	return maior_de_idade
}

func habitantes(slpessoa []Pessoa) map[string]int {
	mapContagem_hab := make(map[string]int)

	for _, v := range slpessoa {
		mapContagem_hab[v.Endereco.Cidade]++
	}
	return mapContagem_hab
}

// ---------MÉTODO DE APRESENTAÇÃO-----------
func (p Pessoa) Formatador() string {
	return fmt.Sprintf("\nOlá, meu nome é %s. Eu tenho %d anos e %.2fm de altura. Moro na cidade %s, na Rua %s, e o número da minha casa é %d",
		p.Nome, p.Idade, p.Altura, p.Endereco.Cidade, p.Endereco.Rua, p.Endereco.Numero)
}

//------------------------------------------

func apresentacao(slpessoa []Pessoa) string {
	var saida_formatada string

	for _, v := range slpessoa {
		saida_formatada += v.Formatador() + "\n"
	}
	return saida_formatada
}

//      NOTAS:

// - 'media_idade()' recebe o slice de Pessoas e retorna a média das idades

// - 'maior_idade()' recebe o slice de Pessoas e 'idade_minima' e retorna
//        apenas aquelas pessoas que tem idade maior ou igual a 18 anos

// - 'habitantes()' recebe o slice de Pessoas e retorna um map onde as chaves
// 	   são os nomes das cidades e os valores são a contagem de pessoas que
// 	   vivem em cada cidade
//
// - 'Apresentacao()' recebe as instâncias da structs Pessoa e retorna uma string
//      com cada pessoa com uma apresentação formatada
