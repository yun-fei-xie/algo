package _6_sweepLine

import "sort"

/*
57. 插入区间
https://leetcode.cn/problems/insert-interval/

给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

方法：扫描线
0.把所有的Interval 放在一起进行排序，然后按照56题的方法做一次区间合并。
1.但是intervals不重叠并且已经按照起始时间进行了排序，但是0的做法显然没有考虑到区间
2.
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	// 下一个的起点和当前的终点进行比较
	// 使用区间生成的方式确定当前待生成区间的起始位置

	var ans = make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		// 当前区间要么是一个新的区间
		if len(ans) == 0 || start > ans[len(ans)-1][1] {
			ans = append(ans, intervals[i])
		} else {
			// 要么融入ans的上一个区间中 (当前区间的start必然是大于或者等于上一个区间的start，因为经过了排序。
			//当前区间的end可能会大于上一个区间的end 也可能小于或者等于。上一个区间的右端点需要更新)
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], end)
		}
	}
	return ans
}

// 考虑到传入的intervals 已经有序 我们只需要合并一段区间范围
// 具体的思路是寻找插入点：

func insert2(intervals [][]int, newInterval []int) [][]int {
	return nil
}
