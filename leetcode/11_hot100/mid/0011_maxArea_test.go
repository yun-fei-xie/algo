package mid

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。


思路：双指针
这个题的双指针比较难想，主要在于双指针的正确性
具体看官方题解：https://leetcode.cn/problems/container-with-most-water/solutions/207215/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/
使用两个指针不断地尝试容器的边界,以获得不同的容器面积

*/

func maxArea(height []int) int {
	max := 0

	left := 0
	right := len(height) - 1

	for left < right {

		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return max
}

func TestMaxArea(t *testing.T) {

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))

}
