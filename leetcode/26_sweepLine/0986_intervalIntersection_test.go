package _6_sweepLine

import (
	"fmt"
	"sort"
	"testing"
)

/*
986. 区间列表的交集
https://leetcode.cn/problems/interval-list-intersections/description/


给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
返回这 两个区间列表的交集 。
形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

这个题和1229的代码模式应该是比较像的

*/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	sort.Slice(firstList, func(i, j int) bool {
		if firstList[i][0] != firstList[j][0] {
			return firstList[i][0] < firstList[j][0]
		} else {
			return firstList[i][1] < firstList[j][0]
		}
	})

	sort.Slice(secondList, func(i, j int) bool {
		if secondList[i][0] != secondList[j][0] {
			return secondList[i][0] < secondList[j][0]
		} else {
			return secondList[i][1] < secondList[j][0]
		}
	})
	var ans = make([][]int, 0)
	l1, l2 := len(firstList), len(secondList)
	for i, j := 0, 0; i < l1 && j < l2; {
		// 分类讨论当前有交集 与没有交集的情况下 i和j如何变化
		fStart, fEnd := firstList[i][0], firstList[i][1]
		sStart, sEnd := secondList[j][0], secondList[j][1]
		// 当前两个 interval没有交集 end小的++
		if fEnd < sStart || fStart > sEnd {
			if fEnd < sEnd {
				i++
			} else {
				j++
			}
			continue
		}

		start := max(fStart, sStart)
		end := min(sEnd, fEnd)
		if end-start >= 0 {
			ans = append(ans, []int{start, end})
			// 当前两个interval有交集 end小的++
			if fEnd < sEnd {
				i++
			} else {
				j++
			}
		}
	}
	return ans

}
func TestIntervalIntersection(t *testing.T) {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
}
