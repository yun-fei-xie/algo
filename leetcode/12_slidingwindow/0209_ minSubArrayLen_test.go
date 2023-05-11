package _2_slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/minimum-size-subarray-sum/
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

这种情况要判断，数组中没有符合条件的子序列
[1,1,1,1,1,1,1,1]
*/
func minSubArrayLen1(target int, nums []int) int {
	left := 0
	right := 0
	min := math.MaxInt64
	sum := 0
	for left <= right && right < len(nums) {
		sum += nums[right]  // right为区间右端点 在本轮循环不动
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

/*
暴力解法：
枚举以right为右端点，以left为左端点的子数字 sum[left ,right] left->[right->0]
*/
func minSubArrayLen2(target int, nums []int) int {
	var ans int = len(nums) + 1
	for right := 0; right < len(nums); right++ {
		var sum int
		// 从right开始，向左扩散
		for left := right; left >= 0; left-- {
			sum += nums[left]
			// 因为区间长度是不断变大的，因此，第一次符合条件的就是以right为右端点的子数组的最小长度
			if sum >= target {
				ans = min(ans, (right-left)+1)
				break
			}
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

/*
枚举以right为区间右端点的连续子数组sum[left..right]
right从0开始，一直枚举到len(nums)-1。
区间左边如何枚举呢？
可以观察到，如果sum[left,right]>=target的，由于数组元素都是正整数（这是一个很重要的条件），所以累加的结果具有单调递增性。
也就是sum[left-1 ,right] > sum[left ,right]
下一轮的区间起始左端点不需要从头开始。而是跟着上一轮的左端点，逐渐缩小。

left什么时候会向右移动？当sum[left ,right] >=target的时候，sum会尝试缩小区间左端点，sum-nums[left]。
如果sum-nums[left] < target的话

这个解法不要看着是两层for循环就觉得是O(n^2)的时间复杂度。
仔细观察就会发现，其实right移动了一轮，left一共也只移动了一轮。这是一个O(n)的时间复杂度。
*/
func minSubArrayLen3(target int, nums []int) int {
	left := 0
	right := 0
	ans := len(nums) + 1
	sum := 0
	for left <= right && right < len(nums) {
		sum += nums[right]             // right为区间右端点 在本轮循环不动
		for sum-nums[left] >= target { //sum先尝试减掉区间的左端点
			sum -= nums[left] //记得把移除掉的left扣掉
			left++
		}
		if sum >= target {
			ans = min(ans, right-left+1)
		}
		right++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res1 := minSubArrayLen2(target, nums)
	fmt.Println(res1)

}
