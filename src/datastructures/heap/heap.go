package heap

import "fmt"

// heap is implemented with a slice

type Heap[T any] struct {
	data []T
	size int
	less func(i, j T) bool
}

func New[T any](less func(i, j int) bool) *Heap[T] {
	h := new(Heap[T])
	h.data = make([]T, 0)
	h.size = 0
	return h
}

func (h *Heap[T]) Push(k T) {
	h.data = append(h.data, k)
	h.size++

	//trickle up
	lastIdx := h.size - 1
	for {
		if lastIdx == 0 {
			break
		}
		p := getParent(lastIdx)
		if h.less(h.data[p], h.data[lastIdx]) {
			h.data[p], h.data[lastIdx] = h.data[lastIdx], h.data[p]
			lastIdx = p
		} else {
			break
		}
	}
}

func (h *Heap[T]) Pop(k T) (T, error) {
	if h.size == 0 {
		var empty T
		return empty, NewHeapEmptyError()
	}
	root := h.data[0]
	lastIdx := h.size - 1
	h.data[0] = h.data[lastIdx]
	h.data = h.data[:h.size-1]
	h.size--

	//trickle down
	lastIdx = 0
	for {
		if lastIdx == h.size-1 {
			break
		}
		lChildIdx := getLeftChild(lastIdx)
		rChildIdx := getRightChild(lastIdx)
		var maxChildIdx int
		if h.less(h.data[lChildIdx], h.data[rChildIdx]) {
			maxChildIdx = rChildIdx
		} else {
			maxChildIdx = lChildIdx
		}
		if h.less(h.data[lastIdx], h.data[maxChildIdx]) {
			h.data[maxChildIdx], h.data[lastIdx] = h.data[lastIdx], h.data[maxChildIdx]
			lastIdx = maxChildIdx
		} else {
			break
		}
	}

	return root, nil
}

func getParent(idx int) int {
	return (idx - 1) / 2
}

func getLeftChild(idx int) int {
	return idx*2 + 1
}

func getRightChild(idx int) int {
	return idx*2 + 2
}

func (h *Heap[T]) Print() {
	if h == nil || len(h.data) == 0 {
		fmt.Println("heap is empty")
		return
	}

	for i := 0; i < len(h.data); i++ {
		fmt.Println(h.data[i])
	}
}
