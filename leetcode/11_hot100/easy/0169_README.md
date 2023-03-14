# 169. 多数元素

## 解题思路



我的解题思路是，直接对数组进行排序。
题目告知了一定会存在超过半数的元素。
因此，排序后的数组在中间位置索引处的元素必然是这个众数。

```


import (
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

```