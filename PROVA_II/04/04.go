package main

//04. Verificar se a pilha contém um valor. Implemente uma função que
//verifica se a pilha contém um valor especificado.
import "fmt"

// Node representa um nó na pilha
type Node struct {
	Value interface{}
	Next  *Node
}

// Stack representa a estrutura de pilha
type Stack struct {
	Top  *Node
	Size int
}

// Push adiciona um elemento ao topo da pilha
func (s *Stack) Push(value interface{}) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
	s.Size++
}

// Pop remove e retorna o elemento do topo da pilha
func (s *Stack) Pop() *interface{} {
	if s.IsEmpty() {
		return nil
	}
	value := s.Top.Value
	s.Top = s.Top.Next
	s.Size--
	return &value
}

// Peek retorna o elemento do topo da pilha sem removê-lo
func (s *Stack) Peek() *interface{} {
	if s.IsEmpty() {
		return nil
	}
	return &s.Top.Value
}

// IsEmpty verifica se a pilha está vazia
func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// Print exibe todos os elementos da pilha
func (s *Stack) Print() {
	current := s.Top
	for current != nil {
		fmt.Println(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	s := &Stack{}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	//vetificar se o elemento '5' está presenta na pilha 'stack'
	if s.verificar(20) {
		fmt.Println("O elemento 5 esta presente na pilha.")
	} else {
		fmt.Println("O elemento 5 nao esta presente na pilha.")
	}
}

func (s *Stack) verificar(value int) bool {
	current := s.Top
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}
