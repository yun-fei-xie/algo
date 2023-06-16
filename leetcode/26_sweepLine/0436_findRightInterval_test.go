package _6_sweepLine

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。注意 i 可能等于 j 。
返回一个由每个区间 i 的 右侧区间 在 intervals 中对应下标组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

方法：二分查找。一个区间[l,r]的右侧区间[k ,p] 必然满足k>=r,并且k是满足条件的所有区间左端点中最小的。(这一步可以使用二分搜索)
也就是说将所有区间的左端点从小到大进行排序，
如何将排序完毕和排序前的位置对应起来。
*/
func findRightInterval(intervals [][]int) []int {

	sortIntervals := make([][]int, len(intervals))
	copy(sortIntervals, intervals)
	// 题目给出每一个区间的起始点都不相同
	sort.Slice(sortIntervals, func(i, j int) bool {
		return sortIntervals[i][0] < sortIntervals[j][0]
	})
	return nil

}

func TestFindRight(t *testing.T) {

	arr := []int{2, 4, 6, 8, 10}
	sort.Ints(arr)
	fmt.Println(sort.SearchInts(arr, 2))

}
