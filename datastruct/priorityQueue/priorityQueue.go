package priorityQueue

import (
	"algo/datastruct/heap"
	"sort"
)

/*
使用堆实现一个优先队列
优先队列本质也是一个队列，它需要实现queue这个接口
支持泛型需要传入一个实现了sort接口的结构 不然在堆中进行上浮和下沉没法比较
*/
type PriorityQueue struct {
	heap *heap.MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	h := heap.NewMaxHeap()
	return &PriorityQueue{heap: h}
}

func (p PriorityQueue) GetSize() int {
	return p.heap.GetSize()
}

func (p PriorityQueue) IsEmpty() bool {
	return p.heap.GetSize() == 0
}

func (p PriorityQueue) EnQueue(item sort.Interface) {

}

func (p PriorityQueue) DeQueue() any {
	return nil
}

func (p PriorityQueue) GetFront() any {
	return nil
}
