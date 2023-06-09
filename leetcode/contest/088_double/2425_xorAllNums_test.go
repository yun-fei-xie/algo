package _88_double_test

import (
	"fmt"
	"testing"
)

/*
2425. 所有数对的异或和
https://leetcode.cn/problems/bitwise-xor-of-all-pairings/description/

给你两个下标从 0 开始的数组 nums1 和 nums2 ，两个数组都只包含非负整数。请你求出另外一个数组 nums3 ，包含 nums1 和 nums2 中 所有数对 的异或和（nums1 中每个整数都跟 nums2 中每个整数 恰好 匹配一次）。
请你返回 nums3 中所有整数的 异或和 。

方法：读题，把nums1中的所有数字和nums2中的所有数字两两组合做一次^操作，然后把所有的结果再做一次^。
[a1,a2] [b1,b2] -> {a1^b1,a1^b2,a2^b1,a2^b2} -> a1^b1^a1^b2^a2^b1^a2^b2
这个时候会发现，如果a1在最后的结果中出现了偶数次，那么a1可以去掉。（a^b^a->b）
什么时候a1会出现偶数次？当nums2的长度是偶数的时候。
不仅如此，当nums2的长度为偶数次的时候，nums1中的所有元素都会出现偶数次。
反之也是一样。

异或：相同为0，不同为1。也就是说，如果a^a->0
*/
func xorAllNums(nums1 []int, nums2 []int) int {
	var ans int

	m1, m2 := len(nums1), len(nums2)
	for i := 0; i < m1; i++ {
		for j := 0; j < m2; j++ {
			ans ^= nums1[i] ^ nums2[j]
		}
	}
	return ans

}

func xorAllNums2(nums1 []int, nums2 []int) int {
	var ans int
	if len(nums2)%2 != 0 {
		for i := 0; i < len(nums1); i++ {
			ans ^= nums1[i]
		}
	}
	if len(nums1)%2 != 0 {
		for i := 0; i < len(nums2); i++ {
			ans ^= nums2[i]
		}
	}

	return ans
}

func TestXor(t *testing.T) {

	fmt.Println(3 ^ 2)
	fmt.Println(3 ^ 2 ^ 3)
}
