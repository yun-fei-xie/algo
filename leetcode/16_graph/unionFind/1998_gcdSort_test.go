package unionFind_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
1998. 数组的最大公因数排序
https://leetcode.cn/problems/gcd-sort-of-an-array/description/

给你一个整数数组 nums ，你可以在 nums 上执行下述操作 任意次 ：
如果 gcd(nums[i], nums[j]) > 1 ，交换 nums[i] 和 nums[j] 的位置。其中 gcd(nums[i], nums[j]) 是 nums[i] 和 nums[j] 的最大公因数。
如果能使用上述交换方式将 nums 按 非递减顺序 排列，返回 true ；否则，返回 false 。

方法：并查集+分解质因数 和0952、6464是同一类题目。
如果nums[i]可以和nums[j]进行交换、nums[j]可以和nums[k]进行交换，那么这三个数字可以进行任意位置的交换。
自然也就可以按照从小到大的次序进行排序。
*/
func gcdSort(nums []int) bool {
	uf := InitUnionFind2(100000 + 1)
	length := len(nums)
	for i := 0; i < length; i++ {
		factors := breakPrimeFactor(nums[i])
		for _, factor := range factors {
			uf.unionElements(factor, nums[i])
		}
	}

	// 合并完成后，对原数组进行排序，然后逐个比较原始数组和当前数组的每一个位置上的数字
	// 如果两者相同，则比较下一位。
	// 如果两者不同，则比较两者是否在同一个集合中。（在同一个集合中就表示可以交换位置）
	numsCopy := make([]int, length)
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	for i := 0; i < length; i++ {
		if nums[i] == numsCopy[i] {
			continue
		} else {
			if uf.find(nums[i]) != uf.find(numsCopy[i]) {
				return false
			}
		}
	}
	return true
}

/*
8 9 4 2 3
2 3 4 8 9
*/
func TestGcdSort(t *testing.T) {
	fmt.Println(gcdSort([]int{8, 9, 4, 2, 3}))
}
