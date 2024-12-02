package main
import (
	"fmt"
)
//Recursão - passe um inteiro e retorne um hexadecimal
func main() {
	var n int

	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	fmt.Println(convHex(n))
}

func convHex(n int) string {
	return fmt.Sprintf("%x", n)
}