package main

import (
	"fmt"

	"github.com/karthikvempati/data-structures-go/bst"
	bst1 "github.com/karthikvempati/data-structures-go/bst"
	heap1 "github.com/karthikvempati/data-structures-go/heap"
	linked_list "github.com/karthikvempati/data-structures-go/linkedlist"
	"github.com/karthikvempati/data-structures-go/queue"
	"github.com/karthikvempati/data-structures-go/stack"
	binary_tree "github.com/karthikvempati/data-structures-go/tree"
)

func main() {
	//testStack()
	//testQueue()
	//testLinkedList()
	//testBinaryTree()
	//testHeap()
	//testMinHeap()
	//testBuildHeap()
	//testDeleteMax()
	//testDeleteMin()
	testBinarySerachTree()
}

func testBinarySerachTree() {
	bst := bst.BinarySearchTree{}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	root := bst.BuildBST(nums)
	array := bst.BstToArray(root)
	fmt.Println(array)
	inorder := bst.InOrderTraversal(root)
	fmt.Println(inorder)
	root = bst1.InsertElement(root, 10)
	fmt.Println(bst.InOrderTraversal(root))
}

func testDeleteMax() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := heap1.Heap{}
	heap.BuildHeap(nums)
	heap.DeleteMaximum()
	heap.PrintHeap()
}

func testDeleteMin() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := heap1.Heap{}
	heap.BuildMinHeap(nums)
	heap.DeleteMinimum()
	heap.PrintHeap()
}

func testHeap() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := heap1.Heap{}
	for i := 0; i < len(nums); i++ {
		heap.InsertElement(nums[i])
	}
	heap.PrintHeap()
}

func testBuildHeap() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := heap1.Heap{}
	heap.BuildHeap(nums)
	heap.PrintHeap()
}

func testMinHeap() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := heap1.Heap{}
	for i := 0; i < len(nums); i++ {
		heap.InsertMinHeapElement(nums[i])
	}
	heap.PrintHeap()
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
