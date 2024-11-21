package main

import "fmt"

func main() {
	x := 10
	var p *int = &x
	fmt.Println(x)

	value := *p
	fmt.Println(value)

	*p = 20

	fmt.Println(x)
	fmt.Println(p)

	//quando uma variável é passada por valor e não por referência,
	//ocorre uma duplicação dessa variável

	w := 100
	z := 200

	var pw *int = &w
	var pz *int = &z

	fmt.Println(w, z)
	change(pw, pz)

	fmt.Println(w, z)

}

//escreva uma função que troque o valor de duas variáveis inteiras usando ponteiros

func change(x *int, y *int) {
	*x, *y = *y, *x
}
