package _3_hashtable

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/4sum/
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// nums[i]+nums[j] 确定，  剩下2个值不确定，降维到3数之和 -> 双指针
			left := j + 1
			right := len(nums) - 1

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left+1] == nums[left] {
						left++
					}
					for left < right && nums[right-1] == nums[right] {
						right--
					}

					left++
					right--
				}

			}

		}
	}
	return res
}

func TestFourSum(t *testing.T) {
	// case1
	//nums := []int{1, 0, -1, 0, -2, 2}
	//target := 0
	//
	//res := fourSum(nums, target)
	//fmt.Println(res)

	// case2
	//nums2 := []int{2, 2, 2, 2}
	//target2 := 8
	//res2 := fourSum(nums2, target2)
	//fmt.Println(res2)

	//case3
	nums3 := []int{-2, -1, -1, 1, 1, 2, 2}
	target3 := 0
	res3 := fourSum(nums3, target3)
	fmt.Println(res3)

}
