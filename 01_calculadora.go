package main

import "fmt"

func main() {
	var escolha int

	for {
		fmt.Printf("\n1 - Adição\n2 - Subtração\n3 - Multiplicação\n4 - Divisão\n5 - Exponenciação\n6 - Radiciação\n7 - Sair")
		fmt.Printf("\n\nEscolha uma opção de 1 a 7: ")
		fmt.Scanln(&escolha)

		if escolha == 7 {
			break
		}
		menu(escolha)
	}
}

func menu(escolha int) int {

	switch escolha {
	case 1:
		fmt.Printf("\nADIÇÃO\n\n")
		soma()
	case 2:
		fmt.Printf("\nSUBTRAÇÃO\n\n")
		subtracao()
	case 3:
		fmt.Printf("\nMULTIPLICAÇÃO\n\n")
		multiplicacao()
	case 4:
		fmt.Printf("\nDIVISÃO\n\n")
		divisao()
	case 5:
		fmt.Printf("\nEXPONENCIAÇÃO\n\n")
		expoente()
	case 6:
		fmt.Printf("\nRADICIAÇÃO\n\n")
		raiz()
	case 7:
		fmt.Printf("\n\nSAINDO...\n")
		//o break é necessário apenas no 'if'
	default:
		fmt.Printf("\nDigite apenas um número de 1 a 5\n")
	}
	return escolha
}

func inputValue() (int, int) {
	var a int
	var b int

	fmt.Printf("Digite o valor de primeiro número: ")
	fmt.Scanln(&a)
	fmt.Printf("Digite o valor do segundo número: ")
	fmt.Scanln(&b)

	return a, b
}

func soma() bool {
	var a int
	var b int

	a, b = inputValue()

	resultado := a + b
	fmt.Printf("\nO resultado da adição é %d\n", resultado)

	return true
}

func subtracao() bool {
	var a int
	var b int

	a, b = inputValue()

	resultado := a - b
	fmt.Printf("\nO resultado da subtração é %d\n", resultado)

	return true
}

func multiplicacao() bool {
	var a int
	var b int

	a, b = inputValue()

	resultado := a * b
	fmt.Printf("\nO resultado da multiplicação é %d\n", resultado)

	return true
}

func divisao() bool {
	var a int
	var b int

	a, b = inputValue()

	resultado := a / b
	resto := a % b
	fmt.Printf("\nO resultado da divisão é %d\n", resultado)
	fmt.Printf("O resto da divisão é %d\n", resto)

	return true
}

func expoente() bool {
	var a int
	var b int

	fmt.Printf("Digite o valor da base: ")
	fmt.Scanln(&a)
	fmt.Printf("Digite o valor do expoente: ")
	fmt.Scanln(&b)

	resultado := a ^ b
	fmt.Printf("\nO resultado da exponenciação é %d\n", resultado)

	return true
}

func raiz() bool {
	var a int
	var b int

	fmt.Printf("Digite o valor da base: ")
	fmt.Scanln(&a)
	fmt.Printf("Digite o valor do índice da raíz: ")
	fmt.Scanln(&b)
	c := 1 / b

	resultado := a ^ c

	fmt.Printf("\nO resultado da exponenciação é %d\n", resultado)

	return true
}
