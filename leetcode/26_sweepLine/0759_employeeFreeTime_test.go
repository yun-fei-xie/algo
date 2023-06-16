package _6_sweepLine

import (
	"sort"
)

/*
759. 员工空闲时间
https://leetcode.cn/problems/employee-free-time/description/

给定员工的 schedule 列表，表示每个员工的工作时间。
每个员工都有一个非重叠的时间段  Intervals 列表，这些时间段已经排好序。
返回表示 所有 员工的 共同，正数长度的空闲时间 的有限时间段的列表，同样需要排好序。

方法：数飞机的变形问题。题目要求找空闲时间，那么可以转化为求解天上没有飞机的时间区间。

	同样将interval拆分为2个点，左端点表示一架飞机起飞，右端点表示一架飞机降落。
	将所有的点按照时间先后进行排序。（**如果两个点坐标相同，那么降落应该在起飞前面**）
*/
type pt struct {
	p int
	f int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	points := make([]*pt, 0, len(schedule)*2)
	for _, sche := range schedule {
		for _, interval := range sche {
			points = append(points, &pt{p: interval.Start, f: 1})
			points = append(points, &pt{p: interval.End, f: -1})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].p != points[j].p {
			return points[i].p < points[j].p
		} else {
			return points[i].f < points[j].f
		}
	})

	var ans = make([]*Interval, 0)
	balance := 0
	for i := 0; i < len(points); i++ {
		point := points[i]
		if balance == 0 && i != 0 {
			if points[i-1].p != point.p { // 同一个点不应该放在最后的结果集合中
				ans = append(ans, &Interval{points[i-1].p, point.p})
			}
		}
		balance += point.f
	}
	return ans
}
