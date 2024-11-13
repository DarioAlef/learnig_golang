package main 
import ("fmt")
//Crie uma função que calcule a soma das diagonais principais e 
//secundárias de uma matriz 3x3 e retorne o resultado
func main() {
	dp := 0
	ds := 0
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{13, 8, 9},
	}
	for i := 0; i < len(matriz); i++ {
		dp += matriz[i][i]    //i = i --> diagonal principal
		ds += matriz[i][2-i]  //i = 2-i --> diagonal secundaria
	}
	fmt.Println("\nSoma da diagonal principal: ", dp)
	fmt.Printf("Soma da diagonal secundaria: %d\n\n", ds)
}