package _1_array

import (
	"fmt"
	"testing"
)

/**
https://leetcode.cn/problems/minimum-size-subarray-sum/
*/

/*
*
滑动窗口
指针i定义为窗口的左边 指针j定义为窗口的右边 初始时，i和j都为0，指向数组的第一个元素
当sum[i:j]<target时，j++
当sum[i:j]>=target时，

	表示得到一个当前最优解，记录之（minLength=j-i+1）
	此时需要在该区间进一步探索更优解
	i ++ (去掉左边一个元素) 计算 sum[i:j] 如果依然符合 sum[i:j]>=target 更新minLength
		如果sum[i:j]<target,j++

什么时候结束搜索：当j>nums.length-1时

看起来还是有点二重循环的感觉

maxLength 定义大一些，不要卡在边界条件上
*/
const maxLength = 1000000

func minSubArrayLen(target int, nums []int) int {
	var i, sum int
	res := maxLength
	subLength := 0

	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLength = j - i + 1
			res = min(subLength, res)
			sum -= nums[i]
			i++
		}
	}
	// check if res was updated
	if res == maxLength {
		return 0
	} else {
		return res
	}
}

func min(i, j int) int {
	if i >= j {
		return j
	}
	return i
}

func TestMinSubArrayLen(t *testing.T) {

	nums := []int{2, 3, 1, 2, 4, 3}
	target := 9

	res := minSubArrayLen(target, nums)
	fmt.Println(res)

}
