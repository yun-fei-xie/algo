package _6_sweepLine

import (
	"sort"
)

/*
数飞机
https://www.lintcode.com/problem/391/

方法：一维扫描线算法
把同一个飞机的起飞和落地拆分为数轴上的两个点(每个点标明是起飞还是降落)（本质也是两个点）
然后对所有的点按照时间的先后循序进行排序，从左到右依次遍历一个点。
如果是起飞点，天上的飞机数量++
如果是降落点，天上的飞机数量--
*/

type Interval struct {
	Start, End int
}

/**
 * @param airplanes: An interval array
 * @return: Count of airplanes are in the sky.
 */
// 将每一个interval拆分为2个点（一个起飞时间点 一个降落时间点）
func CountOfAirplanes(airplanes []*Interval) int {
	// write your code here
	points := make([][]int, 0)
	for _, airplane := range airplanes {
		start := []int{airplane.Start, 1} // 起飞 天上飞机多一架
		end := []int{airplane.End, -1}    // 降落 天上飞机少一架
		points = append(points, start)
		points = append(points, end)
	}
	// 多架次飞机在同一时间起飞降落，那么降落排在前面
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		} else {
			return points[i][1] < points[j][1] // 降落在前
		}
	})

	var ans int = 0
	var cnt int = 0
	for i := 0; i < len(points); i++ {
		if points[i][1] == 1 {
			cnt++
		} else {
			cnt--
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}
