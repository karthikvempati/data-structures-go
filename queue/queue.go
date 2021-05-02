package queue

type Queue struct {
	items []int
}

func (q *Queue) EnQueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) DeQueue() int {
	len := len(q.items)
	val := q.items[0]
	q.items = q.items[1:len]
	return val
}
