package main

import "fmt"

func main() {
	a := 42         //variável normal
	p := &a         //'p' é um ponteiro para a variável 'a'
	fmt.Println(p)  //imprime o endereço de memória 'a'
	fmt.Println(*p) //acessa o valor de 'a' através do ponteiro 'p'

	x := 20
	alterarValor(&x) //passando o endereço de 'x'
	fmt.Println(x)   //resultado = 30 (valor de 'x' foi alterado)

	mainPesssoa()
}
func alterarValor(num *int) {
	*num = *num + 10
}

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

func mainPesssoa() {
	//Criando uma instância pessoa e definindo valores
	pessoa := Pessoa{
		Nome:   "Carlos",
		Idade:  30,
		Cidade: "São Paulo",
	}
	fmt.Println(pessoa)
	fmt.Println("Nome: ", pessoa.Nome)
	pessoa.Idade = 31
	pessoa.Cidade = "Manaus"
	fmt.Println("Nova idade:", pessoa.Idade)
	fmt.Println("Nova cidade: ", pessoa.Cidade)

	//Criando uma instância de Pessoa definindo valores
	//pessoa_endereco:=Pessoa{
	//	Nome:""
	//}
}
func atualizarIdade(p *Pessoa, novaIdade int) {
	p.Idade = novaIdade //Modifica o campo da Idade na struct
}

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
}
type Pessoa_com_Endereco struct {
	Nome   string
	Idade  int
	Cidade string
}
