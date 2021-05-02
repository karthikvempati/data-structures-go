package linkedlist

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) Prepend(val int) {
	c := node{val: val, next: l.head}
	l.head = &c
	l.length++
}
