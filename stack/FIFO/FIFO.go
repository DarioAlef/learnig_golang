package main

import (
	"fmt"
	"math/rand"
)

// Estrutura Node representa o elemento da pilha
type Node struct {
	Value int
	Next  *Node
}

// Estrutura Stack representa a pilha
type Stack struct {
	Top *Node
}

type StackS struct {
	Top *NodeS
}

type NodeS struct {
	Word string
	Next *NodeS
}

// Método Push adiciona um elemento na pilha
func (s *Stack) Push(value int) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
}

func (s *StackS) Put(word string) {
	new_word_node := &NodeS{Word: word, Next: s.Top}
	s.Top = new_word_node
}

// Método Pop remove o elemento do topo da pilha
func (s *Stack) Pop() (int, error) {
	if s.Top == nil {
		return 0, fmt.Errorf("Stack is empty")
	}
	value := s.Top.Value
	s.Top = s.Top.Next
	return value, nil
}

// Método Peek - Retorna o elemento do topo da pilha sem removê-lo
func (s *Stack) Peek() (int, error) {
	if s.Top == nil {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.Top.Value, nil
}

// Método IsEmpety - verifica se a pilha está vazia
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// Método Size - Retorna o tamanho da pilha
func (s *Stack) Size() int {
	size := 0
	current := s.Top

	for current != nil {
		size++
		current = current.Next
	}
	return size
}

// Método Iterate - itera sobre os elementos da pilha
func (s *Stack) Iterate(action func(value int)) {
	current := s.Top
	for current != nil {
		action(current.Value)
		current = current.Next
	}
}

/////////////////////////////////////////////////////////

func main() {
	stack := &Stack{}
	stackS := &StackS{}

	//adicionando elementos na pilha
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stackS.Put("A")
	stackS.Put("B")
	stackS.Put("C")
	stackS.Put("D")

	//iterando sobre os elementos da pilha
	fmt.Println("Iteranting over stack elements: ")
	stack.Iterate(func(value int) {
		fmt.Println(value)
	})

	//mostrando o topo da pilha
	if value, err := stack.Peek(); err == nil {
		fmt.Println("Top element: ", value)
	} else {
		fmt.Println(err)
	}

	//removendo elementos da pilha
	if value, err := stack.Pop(); err == nil {
		fmt.Println("Popped value: ", value)
	} else {
		fmt.Println(err)
	}

	//Verificando se a pilha está vazia
	fmt.Println("Is stack empty? ", stack.IsEmpty())

	//Tamanho da pilha
	fmt.Println("Stack size: ", stack.Size())

	//Verificando o menor elemento da pilha
	if menor, err := stack.Menor(); err == nil {
		fmt.Println("Min element: ", menor)
	} else {
		fmt.Println(err)
	}

	//Verificando o maior elemento da pilha
	if maior, err := stack.Maior(); err == nil {
		fmt.Println("Max element: ", maior)
	} else {
		fmt.Println(err)
	}

	//Invertendo uma string com pilha
	if invertido, err := stackS.Inverter(); err == nil {
		fmt.Println("Inverted string: ", invertido)
	} else {
		fmt.Println(err)
	}
}

////////////////////  EXERCÍCIOS  //////////////////////

// Exercício 1 = Calcular o menor elemento da pilha: Adicione à pilha um método
// chamado 'Min' que retorne o menor elemento armazenado nela.

func (s *Stack) Menor() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Pilha vazia")
	}
	menor := s.Top.Value
	current := s.Top

	for current != nil {
		if current.Value < menor {
			menor = current.Value
		}
		current = current.Next
	}
	return menor, nil
}

// Exercício 2 = Implementar uma pilha de máximo: adicione à pilha um método Max()
//que retorna o maior elemento armazenado nela

func (s *Stack) Maior() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Pilha vazia")
	}

	maior := s.Top.Value
	current := s.Top

	for current != nil {
		if current.Value > maior {
			maior = current.Value
		}
		current = current.Next
	}
	return maior, nil
}

// Exercício 3 = Reverter uma string com ponteiros: use uma pilha implementada para
// criar uma função que inverta uma string. A exemplo, "golang" gera "gnalog"
func (s *StackS) Inverter() (string, error) {

	topo := s.Top
	var reversed string

	for topo != nil {
		reversed += topo.Word
		topo = topo.Next
	}
	//uma pilha consegue "se inverter sozinha", pois o último elemento a entrar será
	//o primeiro a sair, portanto basta retirar da pilha que já estarão invertidos
	return reversed, nil
}

// Exercício 4 = Pilha de Números Aleatórios: crie uma pilha que armazene números aleatórios
//gerados em tempo de execução. Implemente métodos para: Inserir novos números aleatórios;
//Verificar se um número específico está na pilha.

func (s *Stack) aleat() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Pilha vazia")
	}

	numeros := rand.Intn(100)

	return numeros, nil
}
