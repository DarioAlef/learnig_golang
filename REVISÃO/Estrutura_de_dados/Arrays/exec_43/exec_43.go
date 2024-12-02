package main

import "fmt"

func main() {
	// Definindo um array de 5 elementos
	numeros := [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	//Inicializa as variáveis com valor do array para trabalhar no 'for'
	maior := numeros[0]
	segundoMaior := numeros[0]

	//é usado três variáveis, uma para ser o maior valor, a outra para armazenar
	//o que era o segunda maior e a outra para ser de auxiliar (num = auxiliar)
	//(maior = o maior elemento), (segundoMaior = o segundo maior elemento, que
	//que era o valor que pertencia à variável 'maior')
	for _, num := range numeros {
		if num > maior {
			//'segundomaior' recebe o valor que era de maior, mas agora é o segundo maior
			//elemento, e 'maior' recebe o novo elemento de maior valor
			segundoMaior = maior
			maior = num
		} else if num > segundoMaior && num != maior {
			segundoMaior = num
		}
	}

	fmt.Printf("\nO segundo maior número é: %d\n", segundoMaior)
}
