package main

import (
	"bufio"
	"fmt"
	"os"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//declaração explícita com var
	var nome string
	var idade int
	//há um defualt para cada tipo de variável, não apenas 'None'

	//declaração com atribuíção
	var sobrenome string = "Vasconcelos"
	var peso float64 = 85.6

	//atribuição curta (somente dentro da função)
	cor := "pardo"
	temFilhos := true

	println("Nome: ", nome)
	println("Idade: ", idade)
	println("Sobrenome: ", sobrenome)
	fmt.Printf("Peso: %.2f\n", peso)
	println("Cor: ", cor)
	println("Tem filhos?: ", temFilhos)

	//Entrada e saída de dados
	var input1 string
	var input2 string
	var input3 string

	fmt.Printf("Entre com 2 palavras - (Scan): ")
	fmt.Scan(&input1, &input2)
	fmt.Printf("Texto lido - (Scan): %s - %s\n", input1, input2)
	fmt.Scanln(&input3)
	fmt.Printf("Texto digitado - (Scan): %s\n", input3)

	var input4 string
	var input5 int
	fmt.Printf("Digite um texto - (Scanf): ")
	fmt.Printf("%s", &input4)
	fmt.Printf("Texto digitado: %s\n", input4)

	fmt.Printf("Digite um número - (Scanf): ")
	fmt.Printf("%s", &input5)
	fmt.Printf("número digitado - (Scanf): %d", input5)

	//Entrada de Dados - bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Digite um texto com espaços: ")
	texto, _ := reader.ReadString('\n')
	fmt.Printf("Você digitou: %s", texto)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
