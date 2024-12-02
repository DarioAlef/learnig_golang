package main

import "fmt"

//armazene a estrutura de um aluno com seu nome, matrícula e um slice de 4 notas.
//O programa deve permitir CADASTRAR os alunos, LISTAR os alunos com suas notas e
//a MÉDIA, seguida da situação de aprovado ou reprovado, usando '7.0'

type Aluno struct {
	Nome      string
	Matricula int
	Notas     Notas
}
type Notas struct {
	Nota1 float64
	Nota2 float64
	Nota3 float64
	Nota4 float64
}

func main() {
	slalunos := []Aluno{
		Aluno{
			Nome:      "Carlos",
			Matricula: 1234,
			Notas: Notas{
				Nota1: 7.0,
				Nota2: 8.0,
				Nota3: 9.0,
				Nota4: 10.0,
			},
		},
		Aluno{
			Nome:      "Maria",
			Matricula: 5678,
			Notas: Notas{
				Nota1: 5.0,
				Nota2: 6.0,
				Nota3: 7.0,
				Nota4: 8.0,
			},
		},
		Aluno{
			Nome:      "Pedro",
			Matricula: 9012,
			Notas: Notas{
				Nota1: 9.0,
				Nota2: 8.0,
				Nota3: 7.0,
				Nota4: 6.0,
			}}}

	media(slalunos)
}

func media(slalunos []Aluno) {

	for _, v := range slalunos {
		media := (v.Notas.Nota1 + v.Notas.Nota2 + v.Notas.Nota3 + v.Notas.Nota4) / 4
		fmt.Printf("Aluno %s: ", v.Nome)
		if media >= 7.0 {
			fmt.Println("Aprovado")
		} else {
			fmt.Println("Reprovado")
		}
		fmt.Printf("Media Final: %.2f\n\n", media)
	}
}

//func cadastrar() {}

//func listar() {}
