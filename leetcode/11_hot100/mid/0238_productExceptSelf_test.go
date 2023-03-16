package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/product-of-array-except-self/?favorite=2cktkvj

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请不要使用除法，且在 O(n) 时间复杂度内完成此题。

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
时间复杂度要求O(n),并且不能使用除法（先全部相乘，然后挨个做除法 就太简单了 不是考察的点）

题目提示了使用前缀数组和后缀数组进行解题。(前缀积、后缀积)
用一个例子看一下：
nums =   [1 , 2 , 3 , 4 ]
prefix = [1 , 2 , 6 , 24]
postfix = [24,24,12 , 4 ]

可以看到，对于每一个下标i ,  res[i] = prefix[i-1] * postfix[i+1]
很容易理解，例如，对于2这个元素，它等于左边元素乘起来（prefix[i-1]）的结果再乘以右边元素再乘起来（postfix[i+1]）
数组下标越界的问题-> 先判断是否越界，如果越界，返回1

*/

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	mul := 1
	for i := 0; i < len(nums); i++ { // 构造前缀和后缀数组
		prefix[i] = mul * nums[i]
		mul = prefix[i]
	}

	mul = 1
	for i := len(nums) - 1; i >= 0; i-- {
		postfix[i] = nums[i] * mul
		mul = postfix[i]
	}

	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		left := 1
		right := 1
		if i-1 >= 0 {
			left = prefix[i-1]
		}
		if i+1 < len(nums) {
			right = postfix[i+1]
		}

		res[i] = left * right
	}
	return res

}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
