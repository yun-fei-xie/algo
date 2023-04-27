package _1_array

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		} else {
			set[nums[i]] = struct{}{}
		}
	}
	return false
}
