package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/minimum-size-subarray-sum/
这种情况要判断，数组中没有符合条件的子序列
[1,1,1,1,1,1,1,1]
*/
func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	min := math.MaxInt64
	sum := 0
	for left <= right && right < len(nums) {
		sum += nums[right]
		for sum >= target { // 区间合法-> 寻找最小值
			if right-left+1 < min {
				min = right - left + 1
			}
			sum -= nums[left] //记得把移除掉的left扣掉
			left++
		}
		right++
	}
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res1 := minSubArrayLen(target, nums)
	fmt.Println(res1)

}
