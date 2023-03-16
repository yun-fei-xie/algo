package mid

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/merge-intervals/?favorite=2cktkvj
合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。


输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

思路：咋一看不知道如何生成新的区间，但是静下来想想，生成区间就是要确定一个区间的左右边界[left...right]
当我们正在生成一个新的区间，左边界依然确定，就是当前区间的左边界。
那么右边界呢？-> 暂时是当前区间的右边界，然后马上开始考虑下一个区间。
如果下一个区间的左边界<=当前区间的右边界（说明有重叠）则更新当前待生成区间的右边界。
否则的话，一个新的区间生成。

用这种方法去做，需要提前对区间进行排序。


*/

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	if len(intervals) == 0 {
		return res
	}

	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	sort.Slice(intervals, func(i, j int) bool { // 二维slice排序
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})

	for i := 0; i < len(intervals); i++ { // 这里如何控制遍历是一个难点(每次遍历到一个区间，有两种情况 1.作为一个新区间的开始；2.融入到上一个区间中)
		left := intervals[i][0]
		right := intervals[i][1]

		if len(res) == 0 || res[len(res)-1][1] < left { // 作为一个新区间的开始
			res = append(res, []int{left, right})
		} else { // 融合到上一个区间中
			res[len(res)-1][1] = max(res[len(res)-1][1], right)
		}
	}
	return res
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
}
