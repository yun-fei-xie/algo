package _0_double_test

/*
2454. 下一个更大元素 IV

https://leetcode.cn/problems/next-greater-element-iv/description/


给你一个下标从 0 开始的非负整数数组 nums 。对于 nums 中每一个整数，你必须找到对应元素的 第二大 整数。
如果 nums[j] 满足以下条件，那么我们称它为 nums[i] 的 第二大 整数：
j > i
nums[j] > nums[i]
恰好存在 一个 k 满足 i < k < j 且 nums[k] > nums[i] 。
如果不存在 nums[j] ，那么第二大整数为 -1 。
比方说，数组 [1, 2, 4, 3] 中，1 的第二大整数是 4 ，2 的第二大整数是 3 ，3 和 4 的第二大整数是 -1 。
请你返回一个整数数组 answer ，其中 answer[i]是 nums[i] 的第二大整数。

输入：nums = [2,4,0,9,6]
输出：[9,6,6,-1,-1]


方法：单调栈



*/

func secondGreaterElement(nums []int) []int {
	return nums
}
