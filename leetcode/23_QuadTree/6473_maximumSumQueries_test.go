package _3_QuadTree_test

import (
	"testing"
)

/*
6473.最大和查询
https://leetcode.cn/problems/maximum-sum-queries/
给你两个长度为 n 、下标从 0 开始的整数数组 nums1 和 nums2 ，
另给你一个下标从 1 开始的二维数组 queries ，其中 queries[i] = [xi, yi] 。
对于第 i 个查询，在所有满足 nums1[j] >= xi 且 nums2[j] >= yi 的下标 j (0 <= j < n) 中，
找出 nums1[j] + nums2[j] 的 最大值 ，如果不存在满足条件的 j 则返回 -1 。
返回数组 answer ，其中 answer[i] 是第 i 个查询的答案。

思路：
0.题意比较容易理解，对于询问q{k,r} 需要nums1[j]>=k && nums2[j]>=r的提前下，求最大值。
1.最后的答案和j没有关系(不是求j的最大值，比赛的时候理解错误题意。而是求nums1[j]+nums2[j]的最大值)，也就是和元素的顺序没有关系。nums1和nums2中的每一个元素的位置一一对应（类似二维平面的点）

方法：离线查询问题
和之前笔试长亭科技的题目有点像
0. 朴素解法（直接暴力遍历 每次查询，时间复杂度O(n),q次查询 时间复杂度O(qn)）
1. 排序+单调栈+二分
2. 动态开点线段树
3. kd-tree
4. quad-tree
*/

// 使用quad-tree的思路将坐标点进行划分

type quadNode struct {
	topLeft     *quadNode
	topRight    *quadNode
	bottomLeft  *quadNode
	bottomRight *quadNode
	isLeaf      bool
	maxValue    int
}

/*
使用四叉树求解-数组划分使用快速排序的partition的思想
*/
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {

	//xMin := min(nums1...)
	//xMax := max(nums1...)
	//yMin := min(nums2...)
	//yMax := max(nums2...)
	return nil

}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func TestMaximumSumQueries(t *testing.T) {
	//fmt.Println(maximumSumQueries2([]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}))
	//fmt.Println(maximumSumQueries3([]int{3, 2, 5}, []int{2, 3, 4}, [][]int{{4, 4}, {3, 2}, {1, 1}}))
	//fmt.Println(maximumSumQueries3([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 4}}))
	//	fmt.Println(maximumSumQueries3([]int{5, 11, 20, 80, 95, 14, 44, 26, 21, 6, 94, 33, 40, 2, 94, 89},
	//		[]int{60, 76, 61, 6, 7, 71, 22, 26, 100, 63, 17, 2, 89, 19, 100, 69}, [][]int{{80, 18}, {99, 27}, {11, 16}, {36, 86}, {98, 80}, {83, 15}, {47, 31}, {7, 70}, {94, 64}, {16, 48}, {33, 41}, {89, 86}, {61, 54}, {25, 40}, {50, 9}, {38, 84}, {30, 77}, {78, 19}, {88, 15}}))
	//	fmt.Println(maximumSumQueries3([]int{5, 11, 20, 80, 95, 14, 44, 26, 21, 6, 94, 33, 40, 2, 94, 89},
	//		[]int{60, 76, 61, 6, 7, 71, 22, 26, 100, 63, 17, 2, 89, 19, 100, 69}, [][]int{{80, 18}, {99, 27}, {11, 16}, {36, 86}, {98, 80}, {83, 15}, {47, 31}, {7, 70}, {94, 64}, {16, 48}, {33, 41}, {89, 86}, {61, 54}, {25, 40}, {50, 9}, {38, 84}, {30, 77}, {78, 19}, {88, 15}}))
	//
}
