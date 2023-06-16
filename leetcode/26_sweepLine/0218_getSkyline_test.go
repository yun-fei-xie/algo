package _6_sweepLine

import (
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"sort"
	"testing"
)

/*
218.天际线问题
https://leetcode.cn/problems/the-skyline-problem/

城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回 由这些建筑物形成的 天际线 。
每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：
lefti 是第 i 座建筑物左边缘的 x 坐标。
righti 是第 i 座建筑物右边缘的 x 坐标。
heighti 是第 i 座建筑物的高度。
你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。
天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。
注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

方法：用数飞机的方法。把楼的高度看做是不同种类的飞机的架次数量。将每一个点拆分成起飞和降落。
题目要求只标记关键的点。也就是变化的点。

golang这个包实现的优先队列没有Java里面remove一个元素的功能，所以需要另外想办法。
自己实现remove这个功能。
*/

type point struct {
	p     int
	start bool
	high  int
}

// Comparator function (sort by element's priority value in descending order)
func byPriority(a, b interface{}) int {
	priorityA := a.(point).high
	priorityB := b.(point).high
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func getSkyline(buildings [][]int) [][]int {
	queue := pq.NewWith(byPriority) // 优先队列
	points := make([]point, 0)
	for _, building := range buildings {
		points = append(points, point{
			p:     building[0],
			start: true,
			high:  building[2],
		})
		points = append(points, point{
			p:     building[1],
			start: false,
			high:  building[2],
		})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].p != points[j].p {
			return points[i].p < points[j].p
		} else {
			// 如果横坐标相同，高度大的放在前面（这样高度低的不会对结果造成影响 反之不成立）
			return points[i].high > points[j].high
		}
	})
	var ans = make([][]int, 0)
	preHigh := 0
	for i := 0; i < len(points); i++ {
		pt := points[i]
		// 如果p是建筑物的左端点,将这个点入队
		if pt.start == true {
			queue.Enqueue(pt)
		} else {
			// 如果p是建筑物的右端点-> 如果这个点是队列中的最高点，那么最高点出队
			hpoint, _ := queue.Peek()
			if hpoint.(point).high == pt.high {
				queue.Dequeue()
			}

		}
		// 看一下优先队列的最大值
		maxHighPoint, _ := queue.Peek()
		maxHigh := maxHighPoint.(point).high
		if maxHigh != preHigh {
			ans = append(ans, []int{pt.p, maxHigh})
			preHigh = maxHigh
		}
	}
	return ans
}

func TestGetSkyLine(t *testing.T) {
	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}
