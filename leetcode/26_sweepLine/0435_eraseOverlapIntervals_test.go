package _6_sweepLine_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
435. 无重叠区间
https://leetcode.cn/problems/non-overlapping-intervals/description/

给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。
输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

方法：题目要求移除区间最小的数量。
那么还是先对intervals进行排序，如果碰到重叠怎么办？也就是当前区间和上一个区间重叠
经过排序后，后一个区间和当前区间的位置关系只有2种可能：
1.如果当前区间包含上一个区间，排序后区间的位置应该是[4,5]->[4,8] 这个时候应该扔掉当前区间。ps：有没有这种可能-->([4,8]->[4,5]) 首先，经过排序后，不会出现这种情况。
2.如果当前区间overlap上一个区间，{[4,8],[5,7]} -> 这个时候扔掉end更大的区间。
*/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][0]
		}
	})

	// 由于只需要和上一个区间进行对比，因此不需要保存完成的最后的数组
	var pre []int
	var ans int
	for _, interval := range intervals {
		if pre == nil {
			pre = interval
		} else {
			// 不重叠
			if interval[0] >= pre[1] {
				pre = interval
				continue
			} else {
				ans++
				if interval[1] <= pre[1] {
					pre = interval
				}
			}
		}
	}
	return ans
}

func TestEraseOverlap(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}))
}
