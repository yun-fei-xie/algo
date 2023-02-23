package _3_hashtable

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/intersection-of-two-arrays/description/

先把数组转换为Set
然后找出set中的交集 放入到数组中
*/

func intersection(nums1 []int, nums2 []int) []int {

	mapNums1 := make(map[int]struct{})
	mapNums2 := make(map[int]struct{})
	res := make([]int, 0)
	for _, num1 := range nums1 {
		mapNums1[num1] = struct{}{}
	}

	for _, num2 := range nums2 {
		mapNums2[num2] = struct{}{}
	}
	
	for num, _ := range mapNums1 {
		if _, found := mapNums2[num]; found {
			res = append(res, num)
		}
	}
	return res
}

func TestIntersection(t *testing.T) {
	n1 := []int{4, 9, 5}
	n2 := []int{9, 4, 9, 8, 4}

	res := intersection(n1, n2)
	fmt.Println(res)

}
