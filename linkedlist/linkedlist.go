package linkedlist

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(val int) {
	c := Node{Val: val, Next: l.Head}
	l.Head = &c
	l.Length++
}

func (l LinkedList) GetHead() Node {
	return *l.Head
}
