package main

import (
	"fmt"
)

// Estrutura Node representa o elemento da pilha
type Node struct {
	value int
	next  *Node
}

// Estrutura Stack representa a pilha
type Stack struct {
	top *Node
}

// Método Push adiciona um elemento na pilha
func (s *Stack) Push(value int) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
}

// Método Pop remove o elemento do topo da pilha
func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("Stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	return value, nil
}

// Método Peek - Retorna o elemento do topo da pilha sem removê-lo
func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.top.value, nil
}

// Método IsEmpety - verifica se a pilha está vazia
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Método Size - Retorna o tamanho da pilha
func (s *Stack) Size() int {
	size := 0
	current := s.top

	for current != nil {
		size++
		current = current.next
	}
	return size
}

// Método Iterate - itera sobre os elementos da pilha
func (s *Stack) Iterate(action func(value int)) {
	current := s.top
	for current != nil {
		action(current.value)
		current = current.next
	}
}

func main() {
	stack := &Stack{}

	//adicionando elementos na pilha
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

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

	//Verificando saae a pilha está vazia
	fmt.Println("Is stack empty? ", stack.IsEmpty())

	//Tamanho da pilha
	fmt.Println("Stack size: ", stack.Size())

	//Verificando o menor elemento da pilha
	if menor, err := stack.Min(); err == nil {
		fmt.Println("Min element: ", menor)
	} else {
		fmt.Println(err)
	}
}

////////////////////  EXERCÍCIOS  //////////////////////

// Exercício 1 = Calcular o menor elemento da pilha: Adicione à pilha um método
// chamado 'Min' que retorne o menor elemento armazenado nela.

func (s *Stack) Min() (int, error) {
	if s.top == nil { //se o topo tiver vazio
		return 0, fmt.Errorf("Stack is empty") //mensagem reservada para erro
	}
	topo := s.top.value //value puxa o valor do próximo nó
	proximo := s.top.next.value
	var menor int

	if topo > proximo {
		menor = topo
	}

	return menor, nil
}
