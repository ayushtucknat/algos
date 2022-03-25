package heaps

type heaps struct {
	data []int
}

func MaxHeap() *heaps {
	return &heaps{
		data: make([]int, 0),
	}
}

func (h *heaps) Heapify() {}

func (h *heaps) Inset(val int) {
	h.data = append(h.data, val)
	for i := len(h.data) - 1; i > 0; i = i / 2 {
		if h.data[i] <= h.data[i/2] {
			break
		}
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
	}
}

func (h *heaps) IsEmpty() bool {
	return !(len(h.data) > 0)
}

func (h *heaps) Size() int {
	return len(h.data)
}

func (h *heaps) GetMax() *int {
	if len(h.data) != 0 {
		return &h.data[0]
	}
	return nil
}

func (h *heaps) RemoveMax() (max *int) {
	heapSize := len(h.data)
	if heapSize == 0 {
		return
	}
	max = &h.data[0]

	if heapSize == 1 {
		h.data = []int{}
		return
	}
	h.data[0] = h.data[heapSize-1]
	h.data = h.data[:heapSize-1]
	heapSize -= 1
	for i := 0; i < heapSize; {
		if (heapSize >= i*2+1 && h.data[i] > h.data[i*2+1]) && h.data[i] > h.data[i*2+2] {
			break
		}
		if h.data[i*2+2] > h.data[i*2+1] {
			h.data[i], h.data[i*2+2] = h.data[i*2+2], h.data[i]
			i = i*2 + 2
		} else {
			h.data[i], h.data[i*2+1] = h.data[i*2+1], h.data[i]
			i = i*2 + 1
		}
	}
	return
}
