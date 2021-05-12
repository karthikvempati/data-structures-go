package binary_tree

type Queue struct {
	items []*BinaryNode
}

func (q *Queue) EnQueue(item *BinaryNode) {
	q.items = append(q.items, item)
}

func (q *Queue) DeQueue() *BinaryNode {
	len := len(q.items)
	val := q.items[0]
	q.items = q.items[1:len]
	return val
}

func (q *Queue) Peek() *BinaryNode {
	if len(q.items) <= 0 {
		return &BinaryNode{} 
	}
	return q.items[0]
}

func (q *Queue) Clear() {
	q.items = q.items[:0]
}
