package main

import "fmt"

//Crie um map que associe palavras a seus números de ocorrências
// em uma lista de palavras. Imprima o conteúdo do map
func main() {
	var palavra string
	num := 1
	mapPalavras := map[string]int{}

	//mapPalavras[palavra] = num

	fmt.Println("Digite algumas palavra ('0' para parar): ")
	for {
		fmt.Scan(&palavra)
		if palavra == "0" {
			break
		}
		//fmt.Scan(&num)

		for i := range mapPalavras {
			if i == palavra {
				num++
			}
		}
		mapPalavras[palavra] = num
	}

	fmt.Println(mapPalavras)
}
