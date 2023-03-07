# 15. 三数之和

## 题目链接

https://leetcode.cn/problems/3sum/

## 题目描述

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



## 解题思路

外层for循环i，内层滑动窗口left , right

因为数组中有重复的值，所以可以先对数组进行排序，这样相同大小的值会相邻，方便去重逻辑的编写。


## 解题代码

```go

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重left
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				// 去重right
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res


}
```


