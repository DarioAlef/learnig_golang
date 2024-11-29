package main

import (
	"fmt"
)
// 3. Verificar o Primeiro Elemento. Implemente um programa que insere 
//alguns elementos na fila e imprime o primeiro elemento sem removê-lo.

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
	fmt.Println(current)
}

func main() {
	var input int
	var count int
	queue := &Queue{}

	fmt.Println("\nDigite números: ")
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		queue.Enqueue(input)
		count++
	}
	queue.Print()
	fmt.Println("Primeiro elemento: ", *queue.Peek())
}
