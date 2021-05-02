package main

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/stack"
)

func main() {
	testStack()
}

func testStack() {
	s := stack.Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s)
	fmt.Println(s.Pop())
}
