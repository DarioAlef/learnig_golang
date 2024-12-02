package main

import "fmt"

//Contar o Número de Nós: Escreva uma função que conta o número total de nós na árvore

// Node representa um nó da árvore binária
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree representa uma árvore binária
type BinaryTree struct {
	Root *Node
}

// Insert insere um novo valor na árvore binária
func (t *BinaryTree) Insert(value int) {
	newNode := &Node{Value: value}
	if t.Root == nil {
		t.Root = newNode
	} else {
		insertNode(t.Root, newNode)
	}
}

func insertNode(current, newNode *Node) {
	if newNode.Value < current.Value {
		if current.Left == nil {
			current.Left = newNode
		} else {
			insertNode(current.Left, newNode)
		}
	} else {
		if current.Right == nil {
			current.Right = newNode
		} else {
			insertNode(current.Right, newNode)
		}
	}
}

// InOrder faz o percurso in-order na árvore binária
func (t *BinaryTree) InOrder(node *Node, visit func(int)) int {
	count := 0
	if node != nil {
		t.InOrder(node.Left, visit)
		visit(node.Value)
		t.InOrder(node.Right, visit)
		count++
	}
	return count
}

// Search busca por um valor na árvore binária
func (t *BinaryTree) Search(value int) bool {
	return searchNode(t.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	} else if value > node.Value {
		return searchNode(node.Right, value)
	}
	return true
}

// FindMin retorna o menor valor na árvore
func (t *BinaryTree) FindMin() *int {
	current := t.Root
	if current == nil {
		return nil
	}
	for current.Left != nil {
		current = current.Left
	}
	return &current.Value
}

// FindMax retorna o maior valor na árvore
func (t *BinaryTree) FindMax() *int {
	current := t.Root
	if current == nil {
		return nil
	}
	for current.Right != nil {
		current = current.Right
	}
	return &current.Value
}

func main() {
	tree := &BinaryTree{}

	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	for _, value := range values {
		tree.Insert(value)
	}
	fmt.Println("Número de nós:", countNodes(tree.Root))
}

func countNodes(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}
