package _7_monotonicStack

import (
	"container/list"
	"fmt"
	"testing"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	len2 := len(nums2)
	nextGreater := make([]int, len2)

	stack := list.New()
	for j := len2 - 1; j >= 0; j-- {
		for stack.Len() != 0 {
			peek := stack.Back().Value.(int)
			if nums2[j] >= peek {
				stack.Remove(stack.Back())
			} else {
				break
			}
		}

		if stack.Len() == 0 {
			nextGreater[j] = -1
			stack.PushBack(nums2[j])
		} else {
			nextGreater[j] = stack.Back().Value.(int)
			stack.PushBack(nums2[j])
		}
	}

	var ans = make([]int, 0)
	for i := 0; i < len(nums1); i++ {

		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				ans = append(ans, nextGreater[j])
				break
			}
		}
	}
	return ans
}

func TestNextGreaterElement(t *testing.T) {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
