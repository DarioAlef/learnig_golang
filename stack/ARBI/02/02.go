package main

import "fmt"

//Percorrer a Árvore In-Order: Escreva uma função que percorre a
//árvore in-order e imprime os valores em ordem crescente.

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
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(14)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(13)

	//O código executa na árvore uma travessia in-order (esquerda, raiz, direita) começando pela
	//raiz (tree.Root). Para cada nó visitado durante essa travessia, a função anônima é chamada,
	//e o valor do nó (value) é impresso na tela, seguido por um espaço em branco. A função InOrder
	//recebe como parâmetro um nó (node) e uma função (visit) que recebe um valor int, por isso foi
	//necessário passar o tree.Root (também já indica por onde começar) e a função anônima (sem nome
	//e local), o que também faz com que 'value' receba o valor de cada para ser impresso.
	tree.InOrder(tree.Root, func(value int) {
		fmt.Print(value, " ")
	})
}
