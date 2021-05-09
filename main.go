package main

import (
	"fmt"

	linked_list "github.com/karthikvempati/data-structures-go/linkedlist"
	"github.com/karthikvempati/data-structures-go/queue"
	"github.com/karthikvempati/data-structures-go/stack"
	binary_tree "github.com/karthikvempati/data-structures-go/tree"
)

func main() {
	//testStack()
	//testQueue()
	//testLinkedList()
	testBinaryTree()
}

func testBinaryTree() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := binary_tree.CreateBinaryTree(nums)
	//fmt.Println(tree.Root)
	printBinaryTreeLevelOrder(tree.Root)
}

func printBinaryTreeLevelOrder(root *binary_tree.BinaryNode) {
	fmt.Printf("Numbers: %v", root.LevelOrderTraversal())
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
	l := linked_list.LinkedList{}
	l.Prepend(1)
	l.Prepend(5)
	l.Prepend(7)
	l.Prepend(9)
	l.Prepend(3)
	fmt.Println(l)
}
