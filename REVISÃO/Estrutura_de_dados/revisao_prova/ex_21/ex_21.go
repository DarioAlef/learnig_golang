package main
//Crie uma matriz 3x3 de inteiros e escreva uma função que
//imprima cada linha da matriz em uma nova linha no terminal
import "fmt"

func imprimirMatriz(matriz [3][3]int) {

	fmt.Println()

    for i := 0; i < len(matriz); i++ {
        for j := 0; j < len(matriz[i]); j++ {
            fmt.Print(matriz[i][j], " ")
        }
        fmt.Println() // Quebra de linha após imprimir cada linha da matriz
    }
	fmt.Println()
}

func main() {
    var matriz [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
  imprimirMatriz(matriz)
}