package _8_monotonicQueue

import (
	"container/list"
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/sliding-window-maximum/

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7



如果每次都对窗口中的数字遍历肯定会超时。
窗口每次移动，会出去一个元素 进来一个元素。-> 非常像一个入队和出队的操作

如果每次元素进出，这个队列能告诉我当前队列中所有元素的最大值是什么就可以满足需求。

满足单调性的双端队列被称为单调队列。
设计单调队列的时候，pop，和push操作要保持如下规则：
1.pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
2.push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止


*/

type MonotoneQueue struct {
	l *list.List
}

func NewMonotoneQueue() *MonotoneQueue {

	return &MonotoneQueue{l: list.New()}
}

/*
pop操作：让区间左端点和队首元素进行比较，如果队首元素不是value，那么队首元素不需要出栈
*/
func (queue *MonotoneQueue) Pop(value int) {
	if queue.l.Len() != 0 && value == queue.l.Front().Value.(int) {
		queue.l.Remove(queue.l.Front())
	}
}

/*
push操作：让区间右端点val和队尾的元素进行比较，如果value>队尾元素，则移除队尾元素。
直到队列为空，或者是队尾元素大于当前的value值。最后将value放入队尾部。
整个队列从左到右看，呈现出递减的状态。（求区间最大值，队列呈现单调递减）
*/
func (queue *MonotoneQueue) Push(value int) {
	for queue.l.Len() != 0 && value > queue.l.Back().Value.(int) {
		queue.l.Remove(queue.l.Back())
	}
	queue.l.PushBack(value)
}

func (queue *MonotoneQueue) GetMax() int {
	return queue.l.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMonotoneQueue()
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	result := make([]int, 0)
	result = append(result, queue.GetMax())
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k]) // 移除滑动窗口最前面的元素（不一定真的会发生物理删除）
		queue.Push(nums[i])
		result = append(result, queue.GetMax())
	}

	return result
}

func TestMaxSlidingWindow(t *testing.T) {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, -4, 3, 6, 7}, 3))
	//fmt.Println(maxSlidingWindow([]int{1}, 1))
}
