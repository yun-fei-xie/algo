# leetcode 704二分查找


## 题目链接

https://leetcode.cn/problems/binary-search/



## 解题思路

如果不懂二分查找，直接遍历一次数组也能找出答案，时间复杂度为O(n)。  
使用二分查找可以将时间复杂度缩短到log(n)。  
每次从[left , right] 这个区间中寻找与target相等的元素。  
首先取中间位置的索引mid = (left + right) /2 **（题目说明了right + left 的范围不会造成整型溢出）**。
如果target > nums[mid] 则表示target所在的区间在[mid +1 , right] 这个区间。  
如果target < nums[mid] 则表示target所在的区间在[left , mid-1] 这个区间。  
如果target == nums[mid] 则表示找到了值等于target的元素的下标，搜索结束。  

当找到了target 或者整个区间没有任何元素可以被搜索，循环结束。  


## 解题code


```go

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
```