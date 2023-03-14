package mid

import (
	"container/heap"
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/top-k-frequent-elements/description/?favorite=2cktkvj

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

解法1：基于堆排序  golang的container/heap 实现了heap
解法2：基于快速排序

*/

/*
实现sort接口 可以比较
*/
type IHeap [][2]int // 0->数子  1->出现的频率
func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

// 删除数组中最后一个元素
func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	occurences := map[int]int{} // 初始化map
	for _, num := range nums {
		occurences[num]++
	}
	h := &IHeap{}
	heap.Init(h)

	for key, value := range occurences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

func TestTopK(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
