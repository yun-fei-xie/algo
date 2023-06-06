package house_rob_test

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"testing"
)

/*
1696. 跳跃游戏 VI
https://leetcode.cn/problems/jump-game-vi/description/

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。
一开始你在下标 0 处。每一步，你最多可以往前跳 k 步，但你不能跳出数组的边界。也就是说，你可以从下标 i 跳到 [i + 1， min(n - 1, i + k)] 包含 两个端点的任意位置。
你的目标是到达数组最后一个位置（下标为 n - 1 ），你的 得分 为经过的所有数字之和。
请你返回你能得到的 最大得分 。

方法:动态规划+单调队列优化
*/
func maxResult(nums []int, k int) int {

	var dfs func(i int) (score int, flag bool)
	dfs = func(i int) (score int, flag bool) {
		if i >= len(nums) {
			return 0, false
		}
		if i == len(nums)-1 {
			return nums[i], true
		}
		var canReach bool
		m := math.MinInt
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if s, arrive := dfs(j); arrive {
				canReach = true
				m = max(m, s)
			}
		}

		if canReach {
			return nums[i] + m, true
		} else {
			return math.MinInt, false
		}
	}

	maxScore, _ := dfs(0)
	return maxScore
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
动态规划：倒着求解，从最后一个位置出发。由于在最后一个位置上，所以一上来就达到了终点。dp[len(nums)-1] =nums[len(nums)-1]
然后从len(nums)-2 到0倒序枚举每一个位置j。找出dp[j+1 , j+step]这个区间的最大值。
这是正常dp的思路，但是这里找最大值还有一步优化的空间。
找最大值是一个O(k)的操作，乘以dp数组的维度，整个时间的复杂度会来到O(k*N)。能不能说，使用时间复杂度更低的算法，找出[j+1...j+step]的最大值。
单调队列！
*/
func maxResult2(nums []int, k int) int {

	dp := make([]int, len(nums))
	dp[len(nums)-1] = nums[len(nums)-1]
	for j := len(nums) - 2; j >= 0; j-- {
		m := math.MinInt
		for step := 1; step <= k; step++ {
			if j+step < len(nums) {
				m = max(m, dp[j+step])
			} else {
				break
			}
		}
		dp[j] = m + nums[j]
	}
	fmt.Println(dp)
	return dp[0]
}

/*
单调队列优化
*/

func maxResult3(nums []int, k int) int {

	dp := make([]int, len(nums))
	dp[len(nums)-1] = nums[len(nums)-1]
	length := len(nums)
	maxQueue := newMonoQueue()
	maxQueue.enqueue(nums[length-1])
	maxQueue.print()
	for j := length - 2; j >= 0; j-- {
		// [j+1 , j+step]
		if j >= length-k {
			m := maxQueue.extract()
			dp[j] = m + nums[j]
			maxQueue.enqueue(dp[j])

			maxQueue.print()
		} else {
			m := maxQueue.extract()
			dp[j] = m + nums[j]
			maxQueue.dequeue(dp[j+k])
			maxQueue.print()
			maxQueue.enqueue(dp[j])
			maxQueue.print()
		}

	}
	fmt.Println(dp)
	return dp[0]
}

type monoQueue struct {
	l *list.List
}

func newMonoQueue() *monoQueue {
	return &monoQueue{l: list.New()}
}

func (mq *monoQueue) extract() int {
	return mq.l.Front().Value.(int)
}
func (mq *monoQueue) enqueue(val int) {
	for mq.l.Len() != 0 && mq.l.Back().Value.(int) < val {
		mq.l.Remove(mq.l.Back())
	}
	mq.l.PushBack(val)
}
func (mq *monoQueue) dequeue(val int) {
	if mq.l.Len() != 0 && mq.l.Front().Value.(int) == val {
		mq.l.Remove(mq.l.Front())
	}
}
func (mq *monoQueue) print() {
	s := "start<-"
	for front := mq.l.Front(); front != nil; front = front.Next() {
		s += "("
		s += strconv.Itoa(front.Value.(int))
		s += ")" + "<-"
	}
	s += "end"
	fmt.Println(s)
}

func TestMaxResult(t *testing.T) {
	//fmt.Println(maxResult([]int{1, -1, -2, 4, -7, 3}, 2))
	//fmt.Println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3))
	//fmt.Println(maxResult2([]int{1, -1, -2, 4, -7, 3}, 2))
	//fmt.Println(maxResult2([]int{10, -5, -2, 4, 0, 3}, 3))
	//fmt.Println(maxResult3([]int{1, -1, -2, 4, -7, 3}, 2))
	//fmt.Println(maxResult3([]int{10, -5, -2, 4, 0, 3}, 3))
	//fmt.Println(maxResult3([]int{10, -5, -2, 4, 0, 3}, 3))
	//fmt.Println(maxResult2([]int{100, -100, -300, -300, -300, -100, 100}, 4))
	fmt.Println(maxResult3([]int{100, -100, -300, -300, -300, -100, 100}, 4))

}
