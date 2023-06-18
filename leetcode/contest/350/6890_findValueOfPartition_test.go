package _50

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/find-the-value-of-the-partition/


理论最小值：排序后中间两个元素的差
*/

func findValueOfPartition2(nums []int) int {
	sort.Ints(nums)
	var ans int = math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < ans {
			ans = nums[i] - nums[i-1]
		}
	}
	return ans
}

// 正整数
func findValueOfPartition(nums []int) int {
	var ans = math.MaxInt
	var length = len(nums)
	var dfs func(i int, num1Max int, num2Min int, l1 int, l2 int)
	dfs = func(i int, num1Max int, num2Min int, l1 int, l2 int) {
		if i == length {
			return
		}
		if l1 != 0 && l2 != 0 {
			ans = min(ans, absDiff(num1Max, num2Min))
		}
		if nums[i] > num1Max {
			dfs(i+1, nums[i], num2Min, l1+1, l2)
		}

		if nums[i] < num2Min {
			dfs(i+1, num1Max, nums[i], l1, l2+1)
		}

	}

	dfs(0, math.MinInt, math.MaxInt, 0, 0)
	return ans
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func absDiff(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func TestFindValueOfpartition(t *testing.T) {
	fmt.Println(findValueOfPartition([]int{1, 3, 2, 4}))
	fmt.Println(findValueOfPartition([]int{100, 1, 10}))
	fmt.Println(findValueOfPartition([]int{59, 51, 1, 98, 73}))
}
