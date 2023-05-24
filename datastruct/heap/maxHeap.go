package heap

import (
	"fmt"
	"testing"
)

/*
实现一个最大堆

*/

type MaxHeap struct {
	data  []int
	count int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:  make([]int, 0),
		count: 0,
	}
}

func (h *MaxHeap) GetSize() int {
	return h.count
}

/*
对每一个非叶子节点进行一次shiftDown操作 , 倒数第一个非叶子节点的下标是 (count-1)/2
*/
func Heapify(arr []int) *MaxHeap {
	heap := &MaxHeap{
		data:  arr,
		count: len(arr),
	}
	for i := (heap.count - 2) / 2; i >= 0; i-- {
		heap.shiftDown(i)
	}
	return heap
}

func (h *MaxHeap) Insert(item int) {
	h.data = append(h.data, item)
	h.count++
	h.shiftUp(h.count - 1)
}

/*
将最大值和最后一个值交换位置，然后0这个位置的元素进行下沉操作
*/
func (h *MaxHeap) ExtractMax() int {
	max := h.data[0]
	h.data[0], h.data[h.count-1] = h.data[h.count-1], h.data[0] // swap
	h.count--
	h.shiftDown(0)
	return max

}

func (h *MaxHeap) shiftUp(i int) {
	val := h.data[i]
	for i > 0 && val > h.data[i/2] {
		h.data[i] = h.data[i/2]
		i = i / 2
	}
	h.data[i] = val
}

func (h *MaxHeap) shiftDown(i int) {

	for 2*i+1 <= h.count-1 { // 存在左孩子
		j := 2*i + 1
		if j+1 <= h.count-1 && h.data[j+1] > h.data[j] {
			j++
		}
		if h.data[i] > h.data[j] {
			break
		} else {
			h.data[i], h.data[j] = h.data[j], h.data[i]
		}
		i = j
	}
}

func TestMaxHeap(t *testing.T) {
	arr := []int{5, 6, 7, 8}
	heap := NewMaxHeap()
	for i := 0; i < len(arr); i++ {
		heap.Insert(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		max := heap.ExtractMax()
		fmt.Println(max)
	}
}
func TestHeapify(t *testing.T) {
	arr := []int{3, 5, 7, 9, 11, 15, 33, 88, 105, 0}
	heap := Heapify(arr)
	for i := 0; i < len(arr); i++ {
		println(heap.ExtractMax())
	}
}
