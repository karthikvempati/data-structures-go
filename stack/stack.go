package stack

type Stack struct {
	items []int
}

//Push
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

//Push
func (s *Stack) Pop() int {
	len := len(s.items) - 1
	val := s.items[len]
	s.items = s.items[:len-1]
	return val
}
