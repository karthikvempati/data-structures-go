package linkedlist

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(val int) {
	c := Node{val: val, next: l.Head}
	l.Head = &c
	l.Length++
}

func (l LinkedList) GetHead() Node {
	return *l.Head
}
