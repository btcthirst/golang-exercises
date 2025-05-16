package main

import (
	"fmt"
	"stkq/queue"
	"stkq/stack"
)

func main() {
	// Тест стеку
	s := stack.Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Stack:")
	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println(val)
	}

	// Тест черги
	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Queue:")
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		fmt.Println(val)
	}
}
