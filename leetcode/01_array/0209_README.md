# leetcode 209. 长度最小的子数组

## 题目链接

https://leetcode.cn/problems/minimum-size-subarray-sum/

## 题目描述

给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。


## 解题思路

滑动窗口，在窗口滑动的过程中记住最小的窗口值。
窗口如何滑动？（一维数组的区间，用双指针）
用i , j 表示当前的区间端点[i , j] 
如果nums[i] + ... + nums[j] < target , 说明区间元素不够大，此时放入一个元素-> j ++
如果nums[i] + ... + nums[j] >= target , 说明区间元素满足条件，记录下区间长度（j-i+1）并与当前最短长度进行比较。同时从区间移除一个元素（看看能不能让区间更短一点）i++ 
此步骤需要重复进行。

临界条件是：i < len(nums) && j < len(nums)

## 解题代码

```go
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

```
