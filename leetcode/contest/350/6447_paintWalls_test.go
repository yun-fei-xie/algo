package _50

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// 让免费的油漆工人刷cost大的(思路不对)
func paintWalls(cost []int, time []int) int {

	walls := make([][]int, len(cost))
	for i := 0; i < len(cost); i++ {
		walls[i] = []int{cost[i], time[i]}
	}
	sort.Slice(walls, func(i, j int) bool {
		if walls[i][0] != walls[j][0] {
			return walls[i][0] < walls[j][0]
		} else {
			// cost相同，时间长的排在前面，这样可以让免费的工人多刷
			return walls[i][1] < walls[j][1]
		}
	})

	// 双指针
	var ans int
	left, right := 0, len(cost)
	for left < right {

		paid := walls[left]
		ans += paid[0]
		paidTime := paid[1]

		right = right - paidTime
		if right <= left {
			return ans
		}
		left++
	}
	return ans
}

// 转换为0-1背包问题 选或者不选
// 付费+免费=len(cost)= n
// 付费的时间和>= 免费墙的个数 =( n - 付费墙个数)
// time[i0] + time[i1] + time[i3] >= n-3
// (time[i0]+1) + (time[i1]+1) + (time[i3]+1) >= n
// 从time中选择x个元素，使得time+1的累加和>=n 同时，对应的cost和最小
func paintWalls2(cost []int, time []int) int {
	// 任务是将[0...len-1]这个范围的墙刷完
	// i表示当前待考虑的墙的序号 [i...len-1]
	// j表示当前免费可以用的时间
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == len(cost) {
			return 0
		}
		var c1, c2 = math.MaxInt, math.MaxInt
		c1 = dfs(i+1, j+time[i]) + cost[i]
		if j > 0 {
			c2 = dfs(i+1, j-1)
		}
		return min(c1, c2)
	}
	return dfs(0, 0)
	// dfs(i,j)	 = 刷这一面墙 dfs(i+1 , j+time[i]) + cost[i]
	// dfs(i,j) = 不刷这一面墙（用免费额度）dfs(i+1 ， j-1) (前提是j>0)
	// 递归的出口->所有的墙都刷完
	// 上面这个递归函数是错误的。想想一下如果1代表付费0代表免费 那么1001在上面的递归不会出现，但是合法。上面的写法出现了位置依赖。
}

/*
根据paintWalls2修改
*/
func paintWalls3(cost []int, time []int) int {
	var n = len(cost)
	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, 25000)
		for j := 0; j < 25000; j++ {
			mem[i][j] = -1
		}
	}
	// 任务是将[0...len-1]这个范围的墙刷完
	// i表示当前待考虑的墙的序号 [i...len-1] i 的范围[0...n-1]
	// j表示当前免费可以用的时间 j的范围:现在j因为-1的问题可能会出现负数，体现在dfs的时候执行j-1,那么j最多减掉n [-n ,...]
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i == len(cost) {
			if j-n >= 0 {
				return 0
			} else {
				return math.MaxInt32
			}
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		// [i...len-1]这些墙都可以免费刷
		if j-n >= n-i {
			return 0
		}
		mem[i][j] = min(dfs(i+1, j+time[i])+cost[i], dfs(i+1, j-1))

		return mem[i][j]
	}
	return dfs(0, n)
	// dfs(i,j)	 = 刷这一面墙 dfs(i+1 , j+time[i]) + cost[i]
	// dfs(i,j) = 不刷这一面墙（用免费额度）dfs(i+1 ， j-1) (前提是j>0)
	// 递归的出口->所有的墙都刷完
	// 上面这个递归函数是错误的。想想一下如果1代表付费0代表免费 那么1001在上面的递归不会出现，但是合法。上面的写法出现了位置依赖。
}

func TestPaintWalls(t *testing.T) {
	fmt.Println(paintWalls3([]int{2, 3, 4, 2}, []int{1, 1, 1, 1}))
}
