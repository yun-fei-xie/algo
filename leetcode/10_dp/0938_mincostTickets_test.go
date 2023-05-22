package _0_dp

import (
	"fmt"
	"math"
	"testing"
)

/*
938.最低票价
https://leetcode.cn/problems/minimum-cost-for-tickets/?envType=study-plan-v2&envId=dynamic-programming
在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。

火车票有 三种不同的销售方式 ：

一张 为期一天 的通行证售价为 costs[0] 美元；
一张 为期七天 的通行证售价为 costs[1] 美元；
一张 为期三十天 的通行证售价为 costs[2] 美元。
通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张 为期 7 天 的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。

返回 你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费 。



示例 1：

输入：days = [1,4,6,7,8,20], costs = [2,7,15]
输出：11
解释：
例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
在第 1 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 1 天生效。
在第 3 天，你花了 costs[1] = $7 买了一张为期 7 天的通行证，它将在第 3, 4, ..., 9 天生效。
在第 20 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 20 天生效。
你总共花了 $11，并完成了你计划的每一天旅行。
示例 2：

输入：days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
输出：17
解释：
例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
在第 1 天，你花了 costs[2] = $15 买了一张为期 30 天的通行证，它将在第 1, 2, ..., 30 天生效。
在第 31 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 31 天生效。
你总共花了 $17，并完成了你计划的每一天旅行。

方法：
枚举每一次出发前是否买票
1. 如果需要买票，买哪种票
2. 有4种枚举情况（不买票、买种类1、买种类2、买种类3）

如果当前买入某种票，那么可以知道过了多长时间才需要买下一张票

*/

// 这个写法很舒服，但是应该还有优化的空间
func mincostTickets(days []int, costs []int) int {
	// i表示days数组的索引，expired表示上一张票的有效期
	var ans int = math.MaxInt
	var dfs func(i int, expired int, cost int)
	dfs = func(i int, expired int, cost int) {
		if i >= len(days) {
			if cost < ans {
				ans = cost
			}
			return
		}
		// 不需要买票
		if expired >= days[i] {
			dfs(i+1, expired, cost)
		} else {
			// 需要买票
			dfs(i+1, days[i]+0, cost+costs[0])
			dfs(i+1, days[i]+6, cost+costs[1])
			dfs(i+1, days[i]+29, cost+costs[2])
		}
	}
	dfs(0, 0, 0)
	return ans
}

// 修改递归函数 从后向前思考🤔就不用思考前面的状态，因为前面的状态依赖当前的状态
func mincostTickets2(days []int, costs []int) int {
	// days[i...len-1]这个范围买票需要的最小花费
	var dfs func(i int) int
	dfs = func(i int) int {
		// 没有旅行，不需要任何花费
		if i >= len(days) {
			return 0
		}
		// 枚举当前可能需要的花费数量
		// 买costs[0]
		c0 := costs[0] + dfs(i+1)
		// 下一次需要买票的时间
		var j int
		for j = i + 1; j < len(days); j++ {
			if days[j] > days[i]+7-1 {
				break
			}
		}
		c1 := costs[1] + dfs(j)
		// 下一次需要买票的时间
		var k int
		for k = i + 1; k < len(days); k++ {
			if days[k] > days[i]+30-1 {
				break
			}
		}
		c2 := costs[2] + dfs(k)
		return min(c0, c1, c2)
	}
	return dfs(0)
}

/*
记忆化
*/
func mincostTickets3(days []int, costs []int) int {
	// days[i...len-1]这个范围买票需要的最小花费
	var mem = make([]int, len(days))
	for i := 0; i < len(days); i++ {
		mem[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		// 没有旅行，不需要任何花费
		if i >= len(days) {
			return 0
		}
		if mem[i] != -1 {
			return mem[i]
		}

		// 枚举当前可能需要的花费数量
		// 买costs[0]
		c0 := costs[0] + dfs(i+1)
		// 下一次需要买票的时间
		var j int
		for j = i + 1; j < len(days); j++ {
			if days[j] > days[i]+7-1 {
				break
			}
		}
		c1 := costs[1] + dfs(j)
		// 下一次需要买票的时间
		var k int
		for k = i + 1; k < len(days); k++ {
			if days[k] > days[i]+30-1 {
				break
			}
		}
		c2 := costs[2] + dfs(k)
		mem[i] = min(c0, c1, c2)
		return mem[i]
	}
	return dfs(0)
}

func TestMinCostTickets(t *testing.T) {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	fmt.Println(mincostTickets2([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}
