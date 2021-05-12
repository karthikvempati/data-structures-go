package heap

import "fmt"

type Heap struct {
	items     []int
	heap_type string
	count     int
}

func (h *Heap) InsertElement(data int) {
	h.items = append(h.items, data)
	h.count++
	if h.count > 0 {
		h.PercolateUp(h.count - 1)
	}
}

func (h *Heap) InsertMinHeapElement(data int) {
	h.items = append(h.items, data)
	h.count++
	if h.count > 0 {
		h.PercolateUpMinHeap(h.count - 1)
	}
}

func (h *Heap) PrintHeap() {
	fmt.Println(h.items)
}

func (h *Heap) PercolateUp(i int) {
	parent := h.Parent(i)
	for i >= 0 && parent >= 0 && h.items[parent] < h.items[i] {
		temp := h.items[i]
		h.items[i] = h.items[parent]
		h.items[parent] = temp
		i = parent
		parent = h.Parent(i)
	}
}

func (h *Heap) PercolateUpMinHeap(i int) {
	parent := h.Parent(i)
	for i >= 0 && parent >= 0 && h.items[parent] > h.items[i] {
		temp := h.items[i]
		h.items[i] = h.items[parent]
		h.items[parent] = temp
		i = parent
		parent = h.Parent(i)
	}
}

func (h *Heap) PercolateDown(i int) {

}

func (h *Heap) GetMaximum() int {
	return h.items[0]
}

func (h *Heap) DeleteMaximum() int {
	result := h.items[0]
	h.items[0] = h.items[h.count-1]
	h.count--
	return result
}

func (h *Heap) Parent(i int) int {
	if i <= 0 || i >= h.count {
		return -1
	}

	return (i - 1) / 2
}

func (h *Heap) LeftChild(i int) int {
	left := 2*i + 1
	if left >= h.count {
		return -1
	}

	return left
}

func (h *Heap) RightChild(i int) int {
	right := 2*i + 2

	if right >= h.count {
		return -1
	}

	return right
}
