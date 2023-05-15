package _86_double

import (
	"fmt"
	"testing"
)

func findSubarrays(nums []int) bool {

	var prefixSum = 0
	var length = len(nums)
	prefixArr := make([]int, length)
	for i := 0; i < length; i++ {
		prefixArr[i] = prefixSum + nums[i]
		prefixSum = prefixArr[i]
	}

	mem := make(map[int]bool)
	for j := 1; j < length; j++ {
		diff := prefixArr[j] - prefixArr[j-1] + nums[j-1]
		if mem[diff] {
			return true
		} else {
			mem[diff] = true
		}
	}
	return false
}

func TestFindSubArrays(t *testing.T) {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
}
