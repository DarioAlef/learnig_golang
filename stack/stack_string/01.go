package main
import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type Pilha struct {
	Top *Node
}

func (s *Pilha) Push(value string) {
	newNode := &Node{Value: value, Next: s.Top}
	s.Top = newNode
}

func main() {
	stack := &Pilha{}

	stack.Push("A")
	stack.Push("B")
	stack.Push("C")

	if value, err := stack.Inverter(); err == nil {
		fmt.Println("String invertida:", value)
	} else {
		fmt.Println(err)
	}
}

// Exercício 3 = Reverter uma string com ponteiros: use uma pilha implementada para
//criar uma função que inverta uma string. A exemplo, "golang" gera "gnalog"

func (s *Pilha) IsEmpty() bool {
	return s.Top == nil
}

func (s *Pilha) Inverter() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Pilha vazia")
	}

	topo := s.Top
	var reversed string

	for topo != nil {
		reversed += topo.Value
		topo = topo.Next
	}
//uma pilha consegue "se inverter sozinha", pois o último elemento a entrar será
//o primeiro a sair, portanto basta retirar da pilha que já estarão invertidos
	return reversed, nil
}