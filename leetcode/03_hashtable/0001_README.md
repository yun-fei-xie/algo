# leetcode 1 两数之和

## 题目链接

https://leetcode.cn/problems/two-sum/


## 题目描述
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。


## 解题思路

最简单的办法是二重循环暴力求解。
还有一种办法通过递归 dfs 搜索。（这种方法可以求解n数之和）

对二重循环进行改进可以通过hashtable。
先记录下数组中，每个数字的下标


## 解题代码

```go
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

loop:
	for index := 0; index < len(nums); index++ {
		if index2, found := numsMap[target-nums[index]]; found && index != index2 {
			res = append(res, index)
			res = append(res, index2)
			break loop
		}
	}
	return res

}
```