package _6_sweepLine

import (
	"sort"
)

/*
1229. 安排会议日程
https://leetcode.cn/problems/meeting-scheduler/description/

给定两个人的空闲时间表：slots1 和 slots2，以及会议的预计持续时间 duration，请你为他们安排 时间段最早 且合适的会议时间。
如果没有满足要求的会议时间，就请返回一个 空数组。
「空闲时间」的格式是 [start, end]，由开始时间 start 和结束时间 end 组成，表示从 start 开始，到 end 结束。
题目保证数据有效：同一个人的空闲时间不会出现交叠的情况，也就是说，对于同一个人的两个空闲时间 [start1, end1] 和 [start2, end2]，要么 start1 > end2，要么 start2 > end1。





从352过来，下意识想用同样的套路做。先对两个slots按照时间先后进行排序。
然后将slots1放入有序map,以interval的左端为key,右端为val。
然后遍历slots2的每一个interval。然后找交集。
slots2和slots1 中，两个区间相交有3种情况。左边、右边、包含
例如：
左边-> i1:[0,15]与i2:[10,20]
右边-> i1:[15:25]与i2:[10,20]
中间-> i1:[15:20]与i2:[10,20]
*/

/*
暴力求解 击败0%

这里可以发现，某些情况下两个区间肯定没有任何交集。
应该从第一个可能有交集的地方开始查找。
此时应该有二分查找。

另外一种方法是双指针
*/
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {

	sort.Slice(slots1, func(i, j int) bool {
		if slots1[i][0] != slots1[j][0] {
			return slots1[i][0] < slots1[j][0]
		} else {
			return slots1[i][1] < slots1[j][0]
		}
	})

	sort.Slice(slots2, func(i, j int) bool {
		if slots2[i][0] != slots2[j][0] {
			return slots2[i][0] < slots2[j][0]
		} else {
			return slots2[i][1] < slots2[j][0]
		}
	})

	for _, s1 := range slots1 {

		for _, s2 := range slots2 {
			// 讨论位置关系
			s1Left, s1Right := s1[0], s1[1]
			s2Left, s2Right := s2[0], s2[1]
			// 不相交
			if s1Left >= s2Right || s1Right <= s2Left {
				continue
			}
			// 如果相交 如何求出交集的长度？ 将四个点进行排序，然后取出中间两个。就是交集的左右端点。
			points := []int{s1Left, s1Right, s2Left, s2Right}
			sort.Ints(points)
			leftPoint := points[1]
			rightPoint := points[2]
			if rightPoint-leftPoint+1 > duration {
				return []int{leftPoint, rightPoint}
			}
		}
	}
	return nil

}

func minAvailableDuration2(slots1 [][]int, slots2 [][]int, duration int) []int {

	sort.Slice(slots1, func(i, j int) bool {
		if slots1[i][0] != slots1[j][0] {
			return slots1[i][0] < slots1[j][0]
		} else {
			return slots1[i][1] < slots1[j][0]
		}
	})

	sort.Slice(slots2, func(i, j int) bool {
		if slots2[i][0] != slots2[j][0] {
			return slots2[i][0] < slots2[j][0]
		} else {
			return slots2[i][1] < slots2[j][0]
		}
	})
	l1, l2 := len(slots1), len(slots2)
	for i, j := 0, 0; i < l1 && j < l2; {

		// 讨论位置关系
		s1Left, s1Right := slots1[i][0], slots1[i][1]
		s2Left, s2Right := slots2[j][0], slots2[j][1]
		// 找出空闲时间开始的较大值
		start := max(s1Left, s2Left)
		// 找出空闲时间结束的较小值
		end := min(s1Right, s2Right)
		// 这一步求交集长度比上面更加优雅一些
		if end-start+1 > duration {
			return []int{start, start + duration}
		}

		if s1Right < s2Right {
			i++
		} else {
			j++
		}

	}

	return nil

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
func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}
