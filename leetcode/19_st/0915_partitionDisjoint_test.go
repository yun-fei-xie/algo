package _9_st_test

import (
	"fmt"
	"testing"
)

/*
915. 分割数组
https://leetcode.cn/problems/partition-array-into-disjoint-intervals/

给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：
left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回 left 的 长度 。
用例可以保证存在这样的划分方法。
输入：nums = [5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]

方法：st表
1.试想，从左到右遍历数组的下标i，如果能够知道[0...i]这个区间的最大值max1，和[i+1...len(arr))这个区间的最小值min2。
2.如果max1<=min2的话，那么左半部分的所有元素肯定都小于右半部分的所有元素。
3.用两个st表分别维护区间最大值和区间最小值
4.直接使用二维的st表会空间爆炸，由于两种查询都只需要固定一个端点。左边是[0...i] 右边是[i+1，len) 因此，只需要构建1维的表。
*/
func partitionDisjoint(nums []int) int {
	length := len(nums)
	stMax := make([]int, length)
	stMin := make([]int, length)

	stMax[0] = nums[0]
	for i := 1; i < length; i++ {
		stMax[i] = max(nums[i], stMax[i-1])
	}
	stMin[length-1] = nums[length-1]
	for j := length - 2; j >= 0; j-- {
		stMin[j] = min(nums[j], stMin[j+1])
	}

	for k := 0; k < length; k++ {
		leftMax := stMax[k]
		rightMin := stMin[k+1]
		if leftMax <= rightMin {
			return k + 1
		}
	}
	return -1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func TestPartitionDisjoint(t *testing.T) {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
}
