# leetcode 454四数相加

## 题目链接

https://leetcode.cn/problems/4sum-ii/
## 题目描述

给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

## 解题思路

第一反应是深度优先遍历递归求解(枚举)：
组合数量为 n1 * n2 * n3 *n4  100的时候就超时


题解的思路是，知道arr1[i] + arr2[j] 有x种可能等于y。
然后在arr3[k]+arr4[r]中寻找有多少种可能等于-y。

第一次把两个数组的组合相加存入map,map中的keys是两个数组（arr1[i] + arr2[j] ）所有的可能性。


## 解题代码

递归求解

```go

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
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

```


```go
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

```


