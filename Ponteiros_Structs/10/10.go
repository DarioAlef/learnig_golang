package main

import (
	"fmt"
)

//Atualização de Mapas por Ponteiros: escreva uma função incrementarSalario que recebe um
//ponteiro para um map[string]float64 e uma taxa de incremento percentual. A função deve
//atualizar todos os valores no mapa, aumentando o salário de cada pessoa em uma porcentagem.

func main() {
	salarios := map[string]float64{
		"João":  5000.0,
		"Maria": 6000.0,
		"Pedro": 7000.0,
	}

	tx := 0.1
	incrementarSalario(&salarios, tx)
}

func incrementarSalario(m *map[string]float64, tx float64) {
	for k, v := range *m {
		(*m)[k] = v * (1 + tx)
	}
	fmt.Println(*m)
}
