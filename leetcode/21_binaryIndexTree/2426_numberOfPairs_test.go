package _1_binaryIndexTree_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
2426. 满足不等式的数对数目
https://leetcode.cn/problems/number-of-pairs-satisfying-inequality/description/

给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两个数组的大小都为 n ，同时给你一个整数 diff ，统计满足以下条件的 数对 (i, j) ：
0 <= i < j <= n - 1 且
nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.
请你返回满足条件的 数对数目 。

方法：树状数组
1.移项：nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.将相同的索引放在同一侧。
有点高中数列那个感觉了。感觉如果是高中的时候学这些东西，没准对这些信息会更加敏感。nums1[i] - nums2[i] <= nums1[j] - nums2[j] + diff.
令arr[i] = nums1[i]=nums2[i]，于是就有 arr[i] <=arr[j] + diff

2. 经过1的分析，这个问题可以转化为，从右到左遍历，每遇到一个元素arr[i]，需要回答[i+1,len]这个区间有多少个元素满足arr[j]+diff>=arr[i]

3.这个问题和剑指offer51、315CountSmall的不同之处在于多了一个diff。如何处理这个diff呢？
*/
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	length := len(nums1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = nums1[i] - nums2[i]
	}

	set := map[int]struct{}{}
	for _, n := range arr {
		set[n] = struct{}{}
	}
	rank := make([]int, 0)
	for n, _ := range set {
		rank = append(rank, n)
	}
	sort.Ints(rank) // 排序数组 没有重复数据

	bit := newNumBit(length + 1)
	ans := 0
	for i := 0; i < length; i++ {
		idx := sort.SearchInts(rank, arr[i]+diff+1)
		ans += bit.query(idx)
		bit.update(sort.SearchInts(rank, arr[i])+1, 1)
	}
	return int64(ans)
}

type numBit struct {
	parent []int
	n      int
}

func newNumBit(n int) *numBit {
	return &numBit{
		parent: make([]int, n+1),
		n:      n,
	}
}

func (bit *numBit) query(idx int) (ans int) {
	for idx != 0 {
		ans += bit.parent[idx]
		idx -= (-idx) & idx
	}
	return ans
}

func (bit *numBit) update(idx int, val int) {
	for idx <= bit.n {
		bit.parent[idx] += val
		idx += (-idx) & idx
	}
}

func TestNumberOfPairs(t *testing.T) {
	//fmt.Println(numberOfPairs([]int{3, 2, 5}, []int{2, 2, 1}, 1))
	//fmt.Println(numberOfPairs([]int{3, -1}, []int{-2, 2}, -1))

	//arr1: -4,-4, 4,-1,-2,5
	//arr2: -2, 2,-1, 4, 4,3   diff=1
	//arr : 2  -6  5 -5 -6 2
	//sort: -6,-5,2,5
	//res : 0   0  2  1  2 4 -> 9
	//fmt.Println(numberOfPairs([]int{-4, -4, 4, -1, -2, 5}, []int{-2, 2, -1, 4, 4, 3}, 1))

	arr := []int{2, 3, 5, 2, 1}
	sort.Ints(arr)                       // {1,2,2,3,5}
	fmt.Println(sort.SearchInts(arr, 2)) // 重复数据返回第一个位置

}
