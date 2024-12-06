package main

//Intercalar duas filas: crie um programa que intercale os elementos de
//duas filas em uma terceira fila
import "fmt"

type Node1 struct {
	Value interface{}
	Next  *Node1
}

type Queue1 struct {
	Front *Node1
	Rear  *Node1
	Size  int
}

// Enqueue adiciona um elemento ao final da fila
func (q *Queue1) Enqueue(value interface{}) {
	newNode := &Node1{Value: value}
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
func (q *Queue1) Dequeue() *interface{} {
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
func (q *Queue1) Peek() *interface{} {
	if q.IsEmpty() {
		return nil
	}
	return &q.Front.Value
}

// IsEmpty verifica se a fila está vazia
func (q *Queue1) IsEmpty() bool {
	return q.Size == 0
}

// Print exibe todos os elementos da fila
func (q *Queue1) Print() {
	current := q.Front
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

///////////////////////////////////////////////

type Node2 struct {
	Value interface{}
	Next  *Node2
}

type Queue2 struct {
	Front *Node2
	Rear  *Node2
	Size  int
}

// Enqueue adiciona um elemento ao final da fila
func (q *Queue2) Enqueue(value interface{}) {
	newNode := &Node2{Value: value}
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
func (q *Queue2) Dequeue() *interface{} {
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
func (q *Queue2) Peek() *interface{} {
	if q.IsEmpty() {
		return nil
	}
	return &q.Front.Value
}

// IsEmpty verifica se a fila está vazia
func (q *Queue2) IsEmpty() bool {
	return q.Size == 0
}

// Print exibe todos os elementos da fila
func (q *Queue2) Print() {
	current := q.Front
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

/////////////////////////////////////////////

type Node3 struct {
	Value interface{}
	Next  *Node3
}

type Queue3 struct {
	Front *Node3
	Rear  *Node3
	Size  int
}

// Enqueue adiciona um elemento ao final da fila
func (q *Queue3) Enqueue(value interface{}) {
	newNode := &Node3{Value: value}
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
func (q *Queue3) Dequeue() *interface{} {
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
func (q *Queue3) Peek() *interface{} {
	if q.IsEmpty() {
		return nil
	}
	return &q.Front.Value
}

// IsEmpty verifica se a fila está vazia
func (q *Queue3) IsEmpty() bool {
	return q.Size == 0
}

// Print exibe todos os elementos da fila
func (q *Queue3) Print() {
	current := q.Front
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

////////////////////////////////////////////////////////

func main() {
	q1 := &Queue1{}
	q2 := &Queue2{}
	q3 := &Queue3{}

	for i := 0; i < 5; i++ {
		q1.Enqueue(i)
	}

	for i := 5; i < 10; i++ {
		q2.Enqueue(i)
	}

	for q3.Size < 10 {
		q3.Enqueue(*q1.Dequeue())
		q3.Enqueue(*q2.Dequeue())
	}

	q3.Print()
}
