package _8_monotonicQueue_test

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

/*
1438. 绝对差不超过限制的最长连续子数组
https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/\

给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
如果不存在满足条件的子数组，则返回 0 。

方法：如果直到一个窗口中的最大值和最小值，并且最大值和最小值的差值绝对值在limit中，那么这个窗口就是满足条件的。
于是，可以维护两个单调队列。一个maxQueue，一个minQueue。（两个队列可以用同一个结构表示，定义一个比较函数 用户可以自己定义大小）
然后使用双指针维护一个区间。
*/
type comp func(i, j int) int
type monoQueue struct {
	l       *list.List
	compare comp
}

func newMonoQueue(c comp) *monoQueue {
	l := list.New()
	return &monoQueue{
		l:       l,
		compare: c,
	}
}

func (mq *monoQueue) enqueue(val int) {
	for mq.l.Len() != 0 && mq.compare(mq.l.Back().Value.(int), val) > 0 {
		mq.l.Remove(mq.l.Back())
	}
	mq.l.PushBack(val)
}
func (mq *monoQueue) dequeue(val int) {
	if mq.l.Len() != 0 && mq.l.Front().Value.(int) == val {
		mq.l.Remove(mq.l.Front())
	}
}
func (mq *monoQueue) extract() int {
	if mq.l.Len() == 0 {
		return math.MinInt
	}
	return mq.l.Front().Value.(int)
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestSubarray(nums []int, limit int) int {

	maxQueue := newMonoQueue(func(i, j int) int {
		return j - i
	})

	minQueue := newMonoQueue(func(i, j int) int {
		return i - j
	})

	//双指针时间
	ans, left, right := 0, 0, 0
	for right < len(nums) {
		maxQueue.enqueue(nums[right])
		minQueue.enqueue(nums[right])

		minVal := minQueue.extract()
		maxVal := maxQueue.extract()
		if abs(maxVal, minVal) <= limit {
			ans = max(ans, right-left+1)
		} else {
			for left <= right {
				minVal := minQueue.extract()
				maxVal := maxQueue.extract()
				if abs(minVal, maxVal) > limit {
					maxQueue.dequeue(nums[left])
					minQueue.dequeue(nums[left])
					left++
				} else {
					break
				}
			}
		}
		right++
	}
	return ans
}

func TestLongestSubarray(t *testing.T) {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
}
