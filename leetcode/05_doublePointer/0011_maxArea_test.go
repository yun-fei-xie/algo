package _5_doublePointer

/*
https://leetcode.cn/problems/container-with-most-water/description/
11. 盛最多水的容器

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

方法：对向双指针
假如容器的左右边界分别是height[left],height[right]
那么，容器的盛水容量 area = (right-left)*min(height[left] , height[right])（木桶效应）
思考左右边界较低的那一侧，假设是height[left]。
在[left+1...right-1]这个区间中选择任意一个边界作为右边界，求得的面积都要小于area。
这是因为，木桶的整体高度上限已经固定为height[left],同时宽度减小，所以面积肯定会变小。
所以，此时可以将left++ 移动到下一个阶段。
*/
func maxArea(height []int) int {

	var ans int
	var length = len(height)
	for left, right := 0, length-1; left < right; {
		if height[left] < height[right] {
			ans = max(ans, height[left]*(right-left))
			left++
		} else {
			ans = max(ans, height[right]*(right-left))
		}
	}
	return ans
}
