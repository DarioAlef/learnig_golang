package main
//Adicionar Elementos à Fila. Escreva um programa que insere os números de 1 a 10 na fila e imprime todos os elementos
import "fmt"

type Node struct {
    Value interface{}
    Next  *Node
}

type Queue struct {
    Front *Node
    Rear  *Node
    Size  int
}

// Enqueue adiciona um elemento ao final da fila
func (q *Queue) Enqueue(value interface{}) {
    newNode := &Node{Value: value}
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

func main() {
	queue := Queue{}

	for i := 0; i < 10; i++ {
		queue.Enqueue(i + 1)
	}

	queue.Print()

}