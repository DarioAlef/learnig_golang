package main

import "fmt"

//Dado um map que mapeia nomes para idades, crie uma função
//que inverta esse map, ou seja, mapeie idades para nomes
func main() {
	novoMap := make(map[int]string)
	mapPessoas := map[string]int{"Raimundo": 35, "Maria": 25, "Ana": 30, "Paula": 20, "Pedro": 40}

	for nomes, idades := range mapPessoas {
		novoMap[idades] = nomes
	}

	fmt.Println(novoMap)

}
