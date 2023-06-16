package _6_sweepLine

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/meeting-rooms/description/

给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。

示例 1：

输入：intervals = [[0,30],[5,10],[15,20]]
输出：false
示例 2：

输入：intervals = [[7,10],[2,4]]
输出：true


思路：如果会议不冲突，那么就可以全部参加。
如何判断会议不冲突？排序后的区间不重叠

*/

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 || len(intervals) == 1 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool { // 对所有场次，按照开始和结束时间的先后进行排序
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	for i := 1; i < len(intervals); i++ { // 当前场次的开始时间绝对不能小于上一场次的结束时间，
		begin := intervals[i][0]
		if begin < intervals[i-1][1] {
			return false
		}
	}
	return true
}

// 扫描线算法
// 套用391数飞机这道题的模板

func canAttendMeetings2(intervals [][]int) bool {
	points := make([][]int, 0)
	for _, interval := range intervals {
		start := []int{interval[0], 1}
		end := []int{interval[1], -1}
		points = append(points, start)
		points = append(points, end)
	}
	// 如果两个点时间相等，先散会 再开会
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		} else {
			return points[i][1] < points[j][1]
		}
	})
	// 一个人再会的场次最多是1
	var cnt int
	for i := 0; i < len(points); i++ {
		if points[i][1] == 1 {
			cnt++
			if cnt > 1 {
				return false
			}
		} else {
			cnt--
		}
	}
	return true
}

func TestCanAttendMeetings(t *testing.T) {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(canAttendMeetings([][]int{{7, 10}, {2, 4}}))

}
