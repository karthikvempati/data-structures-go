package linkedlist

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	Head   *node
	Length int
}

func (l *LinkedList) Prepend(val int) {
	c := node{val: val, next: l.Head}
	l.Head = &c
	l.Length++
}
