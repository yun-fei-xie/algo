package _88_double_test

import "sort"

/*
方法1：转换成逆序对问题，用一个数据结构查找 <=x的数据的个数。（线段树、树状数组、名次树）
*/
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {

	length := len(nums1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = nums1[i] - nums2[i]
	}
	var ans int
	hashMap := make(map[int]int)
	// 直接使用map这里的时间复杂度太高，需要遍历map
	for i := 0; i < length; i++ {
		for key, val := range hashMap {
			if key <= arr[i]+diff {
				ans += val
			}
		}
		hashMap[arr[i]]++
	}
	return int64(ans)
}

// 树状数组

type BIT []int

// 给包含index=x的每一个元素都加上一个值
func (t BIT) add(x int) {
	for x < len(t) {
		t[x]++
		x += x & -x
	}
}

func (t BIT) query(x int) (res int) {
	for x > 0 {
		res += t[x]
		x &= x - 1
	}
	return
}

func numberOfPair2(a, nums2 []int, diff int) (ans int64) {
	for i, x := range nums2 {
		a[i] -= x
	}
	set := map[int]struct{}{}
	for _, v := range a {
		set[v] = struct{}{}
	}
	b := make(sort.IntSlice, 0, len(set))
	for x := range set {
		b = append(b, x)
	}

	sort.Ints(b)

	t := make(BIT, len(a)+1)
	for _, x := range a {
		ans += int64(t.query(b.Search(x + diff + 1)))
		t.add(b.Search(x) + 1)
	}
	return
}
