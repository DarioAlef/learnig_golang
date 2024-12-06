package main
//Escreva um programa que transfere todos os elementos de uma fila para outra
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

type Queue2 struct {
	Front *Node2
	Rear  *Node2
	Size  int
}

type Node2 struct {
	Value interface{}
	Next  *Node2
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

func (q *Queue2) Print() {
    current := q.Front
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}

func main () {
	q1 := Queue1{}
	q2 := Queue2{}

	for i := 1; i <= 10; i++ {
		q1.Enqueue(i)
	}

	for !q1.IsEmpty() {
		q2.Enqueue(*q1.Dequeue())
	}

	q2.Print()
}