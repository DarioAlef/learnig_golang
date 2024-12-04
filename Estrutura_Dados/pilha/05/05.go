package main
//Verifique se uma string possui parênteses balanceados.

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
	stgr := "((()))()(()()"

	fmt.Printf("\nBalanceado: %v\n\n",isBalancead(stgr))
}

func isBalancead(stgr string) bool {
	stack := &Stack{}

	for _,char := range stgr {
		if char == '(' {
			stack.Push(char)

		} else if char == ')' {
			if stack.IsEmpty() {
				fmt.Println("Nao balanceado")
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}