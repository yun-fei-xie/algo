package _2_slidingwindow

import (
	"fmt"
	"testing"
)

/*
一次遍历,收集遇到的数字1，当碰到0的时候
说明需要进入下一个区间进行统计
如果全是1，走不到else，但是counter会保留结果，因此最后需要通过counter更新一下max
*/
func findMaxConsecutiveOnes1(nums []int) int {
	var ans int
	var counter int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
		} else {
			ans = max(ans, counter)
			counter = 0
		}
	}
	return max(ans, counter)
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	fmt.Println(findMaxConsecutiveOnes1([]int{1, 1, 1, 1, 1}))
}
