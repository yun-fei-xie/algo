package heap

import (
	"fmt"
	"testing"
)

/*
实现一个最小堆
*/
type MinHeap struct {
	data  []int
	count int
}

func NewMinHeap() *MinHeap {

	return &MinHeap{
		data:  make([]int, 0),
		count: 0,
	}

}

func HeapifyMin(arr []int) *MinHeap {

	heap := &MinHeap{
		data:  arr,
		count: len(arr),
	}
	for i := (heap.count - 2) / 2; i >= 0; i-- {
		heap.shiftDown(i)
	}
	return heap
}

func (h *MinHeap) Insert(item int) {
	h.data = append(h.data, item)
	h.count++
	h.shiftUp(h.count - 1)
}

func (h *MinHeap) ExtractMin() int {
	min := h.data[0]
	h.data[0], h.data[h.count-1] = h.data[h.count-1], h.data[0]
	h.count--
	h.shiftDown(0)
	return min

}

func (h *MinHeap) shiftDown(i int) {

	for (2*i + 1) <= h.count-1 { // 存在左孩子

		j := 2*i + 1
		if j+1 <= h.count-1 && h.data[j+1] < h.data[j] {
			j++
		}
		if h.data[i] <= h.data[j] { // 提前终止
			break
		} else {
			h.data[i], h.data[j] = h.data[j], h.data[i]
		}
		i = j
	}
}

func (h *MinHeap) shiftUp(i int) {

	for (i-1)/2 > 0 && h.data[i] < h.data[(i-1)/2] {
		h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
		i = (i - 1) / 2
	}
}

func TestMinHeap(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4}
	heap := HeapifyMin(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(heap.ExtractMin())
	}
}
