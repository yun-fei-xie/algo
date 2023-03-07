package _3_hashtable

import (
	"fmt"
	"testing"
)

/*
代码中的解法体现了一种思想，就是只需要回答限定的问题。
如果能够搜索出全部的组合，那么肯定是能够回答有多少种组合。
但是能够回到有多少种组合，却并不一定能够回答有哪些组合。
https://leetcode.cn/problems/4sum-ii/
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			m[num1+num2]++ // 所有组合计算一次，并统计频数
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += m[-num3-num4] // 同样的模式找匹配
		}
	}
	return count
}

/*
路径爆炸 100 * 100 * 100 * 100 需要搜索一亿次
*/
func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var count = 0
	var mp = make(map[int][]int, 0)
	mp[0] = nums1
	mp[1] = nums2
	mp[2] = nums3
	mp[3] = nums4

	var dfs func(depth int, sum int)
	dfs = func(depth int, sum int) {
		if depth == 4 {
			if sum == 0 {
				count++
			}
			return
		}
		nums := mp[depth]
		for i := 0; i < len(nums); i++ {
			dfs(depth+1, sum+nums[i])
		}
	}
	dfs(0, 0)

	return count
}

func TestForSum(t *testing.T) {
	var arr = []int{-27, 1, -23, -4, -5, -25, -21, 4, -5, 10, 10, 1, -17, -20, -22, -24, -14, 1, -3, -13, -15, -15, -24, -17, -31, -12, -14, -16, -29, -4, -16, -26, -5, 7, 1, -21, -18, -24, 1, -27, -31, -12, 1, 5, -13, 10, -28, -1, -4, -2, 10, -3, -20, -29, -7, -14, -21, -31, -1, -24, -5, -8, -2, 5, -7, -26, -27, -19, 8, -25, -25, -1, -20, -9, -13, -22, -10, -4, -14, -20, -10, 5, -23, -3, -17, -3, -8, -16, -10, -29, -24, 9, -11, 8, -4, -5, -11, -27, -25, -26}
	fmt.Println(len(arr))

}
