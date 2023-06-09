package _1_binaryIndexTree_test

import "sort"

/*
315. 计算右侧小于当前元素的个数
https://leetcode.cn/problems/count-of-smaller-numbers-after-self/description/
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

方法：树状数组
这个题目和剑指offer第51题是一模一样。
*/
func countSmaller(nums []int) []int {
	length := len(nums)
	var rank = make([]int, length)
	copy(rank, nums)
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		rank[i] = sort.SearchInts(nums, rank[i]) + 1 // 排名
	}
	var ans = make([]int, length)
	bitree := newBinaryIndexTree(length)
	for i := length - 1; i >= 0; i-- {
		cnt := bitree.query(rank[i] - 1)
		ans[i] = cnt
		bitree.update(1, rank[i])
	}
	return ans
}

type binaryIndexTree struct {
	parent []int
	n      int
}

func newBinaryIndexTree(n int) *binaryIndexTree {
	p := make([]int, n+1)
	return &binaryIndexTree{
		parent: p,
		n:      n,
	}
}

func (bit *binaryIndexTree) query(idx int) (ans int) {
	for idx != 0 {
		ans += bit.parent[idx]
		idx -= idx & (-idx)
	}
	return ans
}
func (bit *binaryIndexTree) update(val int, idx int) {
	for idx <= bit.n {
		bit.parent[idx] += val
		idx += idx & (-idx)
	}
}
