package main
//Implemente um programa que inverte a ordem dos elementos de uma fila usando uma pilha auxiliar
import "fmt"

type Nodef struct {
    Value interface{}
    Next  *Nodef
}

type Queue struct {
    Front *Nodef
    Rear  *Nodef
    Size  int
}

// Enqueue adiciona um elemento ao final da fila
func (q *Queue) Enqueue(value interface{}) {
    newNode := &Nodef{Value: value}
    if q.Rear == nil {
        q.Front = newNode
        q.Rear = newNode
    } else {
        q.Rear.Next = newNode
        q.Rear = newNode
    }
    q.Size++
}

// Dequeue remove e retorna o elemento no início da fila
func (q *Queue) Dequeue() *interface{} {
    if q.IsEmpty() {
        return nil
    }
    value := q.Front.Value
    q.Front = q.Front.Next
    if q.Front == nil {
        q.Rear = nil
    }
    q.Size--
    return &value
}

// Peek retorna o elemento no início da fila sem removê-lo
func (q *Queue) Peek() *interface{} {
    if q.IsEmpty() {
        return nil
    }
    return &q.Front.Value
}

// IsEmpty verifica se a fila está vazia
func (q *Queue) IsEmpty() bool {
    return q.Size == 0
}

// Print exibe todos os elementos da fila
func (q *Queue) Print() {
    current := q.Front
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}

///////////////////////////////////////////////////////////////////

type Nodep struct {
	Value interface{}
	Next  *Nodep
}

// Stack representa a estrutura de pilha
type Stack struct {
	Top  *Nodep
	Size int
}

// Push adiciona um elemento ao topo da pilha
func (s *Stack) Push(value interface{}) {
	newNode := &Nodep{Value: value, Next: s.Top}
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
	q := &Queue{}
	s := &Stack{}

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	for !q.IsEmpty() {
		s.Push(*q.Dequeue())
	}

	for !s.IsEmpty() {
		q.Enqueue(*s.Pop())
	}

	q.Print()
}
