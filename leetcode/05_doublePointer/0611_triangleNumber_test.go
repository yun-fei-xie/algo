package _5_doublePointer

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/valid-triangle-number/
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。


三角形边长的定义：任何两边之和大于第三边
a + b > c
a + c > b
b + c > a
这是没有边长大小关系的约束。
如果已经知道了三条边种，a最长。
a + b > c 和 a + c > b 恒成立，不需要做进一步的验证。
只需要验证是否 b + c > a

现在问题转化为，如何求一个排序数组中满足nums[i] + nums[j] > target的元素对数。
如果nums[i] + nums[j] > target的话，那么 nums[i+1]、nums[i+2]...nums[j-1] 和nums[j]相加已经也是。
这里就出现了（j-i）对合法的元素。更新结果， ans+=(j-i) , 更新j--。
如果nums[i] + nums[j] <= target的话，说明nums[i]不够大，i++。
可以发现，这样做其实是从外层循环到内层循环，每次固定一条较大边。

*/

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var length = len(nums)
	var ans int
	for i := 2; i < length; i++ {
		for left, right := 0, i-1; left < right; {
			sum := nums[left] + nums[right]
			if sum > nums[i] {
				ans += right - left
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
func TestTriangleNumber(t *testing.T) {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}
