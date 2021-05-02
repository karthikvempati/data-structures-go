package main

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/linkedlist"
	"github.com/karthikvempati/data-structures-go/queue"
	"github.com/karthikvempati/data-structures-go/stack"
)

func main() {
	//testStack()
	// testQueue()
	testLinkedList()
}

func testStack() {
	s := stack.Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s)
	fmt.Println(s.Pop())
}

func testQueue() {
	q := queue.Queue{}
	q.EnQueue(100)
	q.EnQueue(200)
	q.EnQueue(300)
	fmt.Println(q)
	fmt.Println(q.DeQueue())
	fmt.Println(q)
}

func testLinkedList() {
	l := linkedlist.LinkedList{}
	l.Prepend(1)
	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)
	l.Prepend(3)
	fmt.Println(l)
}
