package _2_slidingwindow

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/subarray-product-less-than-k/description/

给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
示例 1：
输入：nums = [10,5,2,6], k = 100
输出：8
解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
示例 2：
输入：nums = [1,2,3], k = 0
输出：0

区间中的元素都是正数(思考，如果没有这个条件限制，应该怎么做)

*/

/*
方法1：双指针。和209题的思想基本一致。
枚举以right为区间右端点的子数组。如果nums[left ,right] 这段区间的乘积小于k
那么nums[left+1 ,right]、nums[left+2 ,right] ...、nums[right ,right] 的乘积都小于k
这一段，一共有(right-left+1)个子数组符合条件。 也就是区间的长度

如果k<=1 肯定无解。因为每一个元素都是正整数。nums[i] 必然大于1。
而且mul初始化就是1

需要注意的是，区间的长度可能为0。
想像这样一个数组[4,5,6] k = 3
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	var ans int
	var mul = 1

	for right, left := 0, 0; right < len(nums); right++ {
		mul *= nums[right]

		for mul >= k {
			mul = mul / nums[left]
			left++
		}
		ans += right - left + 1

	}
	return ans
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
