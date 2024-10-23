package main

type heapNode struct {
	data        string
	freq        int
	left, right *heapNode
}

type minHeap struct {
	size int
	data []*heapNode
}

func buildMinHeap(a []*heapNode) *minHeap {
	heapSize := len(a)

	for i := (heapSize - 1) / 2; i >= 0; i-- {
		minHeapify(a, heapSize, i)
	}

	var heap *minHeap = &minHeap{size: heapSize, data: a}
	return heap
}

func minHeapify(a []*heapNode, size int, i int) {
	left := (i * 2) + 1
	right := (i * 2) + 2

	minimum := i

	if left < size && a[left].freq < a[i].freq {
		minimum = left
	}

	if right < size && a[right].freq < a[minimum].freq {
		minimum = right
	}

	if minimum != i {
		a[i], a[minimum] = a[minimum], a[i]
		minHeapify(a, size, minimum)
	}
}

func popMinHeap(h *minHeap) *heapNode {
	a := h.data
	temp := a[0]

	a[0] = a[h.size-1]

	h.size--
	minHeapify(a, h.size, 0)
	return temp
}

func pushMinHeap(h *minHeap, n *heapNode) {
	h.size++
	i := h.size - 1
	h.data = append(h.data, n)

	for i != 0 && n.freq < h.data[(i-1)/2].freq {
		h.data[i] = h.data[(i-1)/2]
		i = (i - 1) / 2
	}

	h.data[i] = n
}
