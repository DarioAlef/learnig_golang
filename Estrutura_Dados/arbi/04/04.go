package main
//Escreva um programa que encontra e imprime o menor valor na árvore binária
import "fmt"

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

	for i := 0; i < 10; i++ {
		tree.Insert(i + 1)
	}

	fmt.Println(*tree.FindMin())
}