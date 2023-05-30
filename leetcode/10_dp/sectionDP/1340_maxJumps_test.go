package sectionDP_test

import (
	"fmt"
	"testing"
)

/*
1340. 跳跃游戏 V
https://leetcode.cn/problems/jump-game-v/description/

给你一个整数数组 arr 和一个整数 d 。每一步你可以从下标 i 跳到：
i + x ，其中 i + x < arr.length 且 0 < x <= d 。
i - x ，其中 i - x >= 0 且 0 < x <= d 。
除此以外，你从下标 i 跳到下标 j 需要满足：arr[i] > arr[j] 且 arr[i] > arr[k] ，其中下标 k 是所有 i 到 j 之间的数字（更正式的，min(i, j) < k < max(i, j)）。
你可以选择数组的任意下标开始跳跃。请你返回你 最多 可以访问多少个下标。
请注意，任何时刻你都不能跳到数组的外面。

方法：递归+记忆化搜索 感觉也不是很难
1.站在一个点i，可以向左跳到j，也可以向右到j。（如果符合条件）
2.什么时候站在一个点不能跳了？这个点的左右都比它高。（base case）
3.由于站在点i，必须要找到跳跃点数最大的路径，那么，需要枚举向左或者向右可能的跳法，取最大值。
*/
func maxJumps(arr []int, d int) int {
	var mem = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		mem[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		// 记忆化
		if mem[i] != -1 {
			return mem[i]
		}
		var ret int = 1
		defer func() {
			mem[i] = ret
		}()
		// 对于边界的处理
		if i == 0 && i+1 < len(arr) && arr[i] <= arr[i+1] {
			return 1
		} else if i == len(arr)-1 && i-1 >= 0 && arr[i] <= arr[i-1] {
			return 1
		} else if i > 0 && i < len(arr)-1 && arr[i] <= arr[i-1] && arr[i] <= arr[i+1] {
			return 1
		}

		// 向右跳 只要有遮挡就跳不过去 碰到遮挡直接退出循环
		for j := i + 1; j < len(arr) && j <= i+d; j++ {
			if arr[i] > arr[j] {
				ret = max(ret, dfs(j)+1)
			} else {
				break
			}
		}
		// 向左跳
		for j := i - 1; j >= 0 && j >= i-d; j-- {
			if arr[i] > arr[j] {
				ret = max(ret, dfs(j)+1)
			} else {
				break
			}
		}
		return ret
	}

	var ans = 1
	for i := 0; i < len(arr); i++ {
		ans = max(dfs(i), ans)
	}
	return ans
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if m < args[i] {
			m = args[i]
		}
	}
	return m
}

func TestMaxJump(t *testing.T) {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	fmt.Println(maxJumps([]int{3, 3, 3, 3, 3}, 2))
	fmt.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
	fmt.Println(maxJumps([]int{66}, 1))
}
