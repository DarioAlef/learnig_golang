package main

import "fmt"

//Buscar um Valor na Árvore: Implemente um programa que verifica
// se o valor 6 está presente na árvore

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
func (t *BinaryTree) InOrder(node *Node, visit func(int)) {
	if node != nil {
		t.InOrder(node.Left, visit)
		visit(node.Value)
		t.InOrder(node.Right, visit)
	}
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
	target := 6

	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(1)
	//tree.Insert(6)

	if tree.Search(target) {
		fmt.Println("O valor", target, "está na árvore.")
	} else {
		fmt.Println("O valor", target, "não está na árvore.")
	}
}
