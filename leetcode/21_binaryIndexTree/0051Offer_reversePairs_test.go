package _1_binaryIndexTree_test

import (
	"sort"
)

/*
数组中逆序对的个数
方法：树状数组+离散化 先看非离散化的做法
1.arr=[5,5,2,3,6]，用一个桶记录每个元素出现的次数bucket=[0 1 1 0 2 1 0 0 0]
2.假设当前遍历到arr[i]，那么通过查找bucket[i+1:len]这段区间的后缀和，便可以知道arr中，下标大于i的元素有多少个。
3.简单这样做肯定有问题，arr的值可能有负数，数组的下标没有负数。（可以通过平移解决）;arr的值域很大，但是可能会很离散。
3.总的来说：可以通过从右向左进行遍历

离散化：离散化一个序列的前提是我们只关心这个序列里面元素的相对大小，而不关心绝对大小（即只关心元素在序列中的排名）
*/

func reversePairs(nums []int) int {

	// 排序是为了二分，进而拿到每个元素的名次
	rank := make([]int, len(nums))
	// copy一份是为了二分的时候知道传入的数据在整个数组中的排名。本题不只是需要知道排名，还要考虑位置。
	copy(rank, nums)

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		rank[i] = sort.SearchInts(nums, rank[i]) + 1 // 名次从1开始计算 为了适应BIT的结构
	}
	b := newBit(len(rank))
	var ans int
	for i := len(nums) - 1; i >= 0; i-- {

		r := rank[i]
		ans += b.query(r - 1)
		b.update(r)
	}
	return ans
}

/*
bit中数组的每一个下标表示一个名词
*/
type bit struct {
	parent []int
	n      int
}

func newBit(n int) *bit {
	parent := make([]int, n+1)
	return &bit{
		parent: parent,
		n:      n,
	}
}

/*
查找[1:idx]闭区间的区间和 也就是前缀和
*/
func (b *bit) query(idx int) (ans int) {
	for idx != 0 {
		ans += b.parent[idx]
		idx -= idx & (-idx)
	}
	return ans
}

/*
将idx这个名次拥有的元素数量++
*/
func (b *bit) update(idx int) {
	for idx <= b.n {
		b.parent[idx]++
		idx += idx & (-idx)
	}
}
