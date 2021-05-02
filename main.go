package main

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/queue"
	"github.com/karthikvempati/data-structures-go/stack"
)

func main() {
	//testStack()
	testQueue()
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
