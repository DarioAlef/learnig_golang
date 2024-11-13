package main 
import ("fmt")
//Crie um map onde as chaves são IDs (inteiros) e os valores são structs Pessoa 
//(com campos nome e idade). Adicione algumas pessoas ao map e imprima os dados

type Pessoa struct {
	Nome string
	Idade int
}
func main() {
	//var ID int
	sl := []Pessoa{
		{
			Nome:  "Maria",
			Idade: 25,
		},
		{
			Nome:  "Pedro",
			Idade: 30,
		},
		{
			Nome:  "Joaquim",
			Idade: 35,
		},
	}
	mapPessoas := make(map[int]Pessoa)

	for id, value := range sl {
		mapPessoas[id] = value
	}

	fmt.Println(mapPessoas)	
}