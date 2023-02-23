package _3_hashtable

import (
	"fmt"
	"sort"
	"testing"
)

/*
去重的逻辑
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重left
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				// 去重right
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res

}

func TestThreeSum(t *testing.T) {

	nums := []int{-1, 0, 1, 2, -1, -4} // {-1 , -1 , 0 , 1, 2, 4}  这个用例会有重复
	nums2 := []int{-2, 0, 0, 2, 2}
	res := threeSum(nums)
	res2 := threeSum(nums2)
	fmt.Println(res)
	fmt.Println(res2)

}
