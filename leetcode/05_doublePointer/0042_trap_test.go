package _5_doublePointer

/*
https://leetcode.cn/problems/trapping-rain-water/
42.接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

方法1，前缀数组、后缀数组
单独思考每一个柱子能够接到的雨水的面积（二维）。
对于柱子i来说，从左到右height[0...i]的前缀最大值preMax[i]、从右到左height[i...len-1]的后缀最大值suffixMax[i]
min(preMax[i], suffixMax[i])构成了当前柱子的雨水最大高度。
于是当前柱子能够接到的雨水就是 min(preMax[i], suffixMax[i]) - height[i]

可以将height[i]认为是当前桶里面已经被占用了多少空间。
height   : 0,1,0,2,1,0,1,3,2,1,2,1
preMax   : 0,1,1,2,2,2,2,3,3,3,3,3
suffixMax: 3,3,3,3,3,3,3,3,2,2,2,1
接水      : 0,0,1,0,1,2,1,0,0,1,0,0  -> 6

时间复杂度O（n） 空间复杂度O（n）

方法2，前缀数组、后缀数组、对向双指针
方法1中使用了前缀数组、后缀数组。
思考这样一种情况，假设对于索引i的位置，preMax[i]已经知道，但是suffixMax[i]不知道，
但是此时已经知道了suffixMax[j] (j>i)，并且已经知道 preMax[i]<=suffixMax[j]
那么柱子的最大的接雨水的高度就是preMax[i]。 这是因为 suffixMax[i] 必然是 >= suffixMax[j]
于是可以用对向双指针的方法，边遍历边更新。

*/

func trap1(height []int) int {
	var maxHeight = 0
	var length = len(height)
	preMax := make([]int, length)
	suffixMax := make([]int, length)
	for i := 0; i < length; i++ {
		preMax[i] = max(maxHeight, height[i])
		maxHeight = preMax[i]
	}

	maxHeight = 0
	for j := length - 1; j >= 0; j-- {
		suffixMax[j] = max(maxHeight, height[j])
		maxHeight = suffixMax[j]
	}

	var ans int

	for k := 0; k < length; k++ {
		ans += min(preMax[k], suffixMax[k]) - height[k]
	}
	return ans
}

func trap2(height []int) int {
	var preMax = 0
	var suffixMax = 0
	var ans = 0
	var length = len(height)

	for left, right := 0, length-1; left <= right; {
		preMax = max(preMax, height[left])
		suffixMax = max(suffixMax, height[right])

		if preMax <= suffixMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += suffixMax - height[right]
			right--
		}

	}
	return ans
}
