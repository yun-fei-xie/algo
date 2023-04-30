package _0_dp

import (
	"fmt"
	"testing"
)

/*
没有思路的时候尝试枚举
1.从起点出发，每次选择一个和当前点不一样的点，走过去。
2.在油料消耗完之前，查看是否到达了终点。如果达到了终点，线路++
3.油量就是递归终止条件

穷举的思路是对的，但是遇到油量太多的时候，需要大量递归，时间复杂度太高
使用记忆化搜索，需要使用后序遍历的方式
*/
func countRoutes(locations []int, start int, finish int, fuel int) int {

	var ans int
	var dfs func(current int, remainFuel int)
	dfs = func(current int, remainFuel int) {
		// 不需要写递归返回的语句，因为流量不够，不会进入递归
		if current == finish {
			ans++
		}
		// 每一轮都可以选择除了出发点外的所有点
		for i := (current + 1) % len(locations); i != current; i = (i + 1) % len(locations) {
			// 可以去成
			costFuel := cost(locations[current], locations[i])
			if costFuel <= remainFuel {
				dfs(i, remainFuel-costFuel)
			}
		}
	}
	dfs(start, fuel)
	return ans
}

/*
递归+记忆化搜索
*/
func countRoutes2(locations []int, start int, finish int, fuel int) int {
	// mem[i][j] 表示从位置i出发，油量为j的情况下到达finish的路线数量。初始为-1，区别于数量0。
	var mod = 1000000007
	mem := make([][]int, len(locations))
	for i := 0; i < len(mem); i++ {
		fl := make([]int, fuel+1)
		for j := 0; j < fuel+1; j++ {
			fl[j] = -1
		}
		mem[i] = fl
	}

	var traceback func(current int, finish int, remainFuel int) int
	traceback = func(current int, finish int, remainFuel int) int {
		// 缓存中存在答案
		if mem[current][remainFuel] != -1 {
			return mem[current][remainFuel]
		}
		// 如果当前油量为0 并且当前的位置不是终点
		if remainFuel == 0 && current != finish {
			mem[current][0] = 0
			return 0
		}
		// 油量不为0，但是无法到达任何位置
		var hasNext bool
		for i := 0; i < len(locations); i++ {
			if i != current {
				if remainFuel >= cost(locations[current], locations[i]) {
					hasNext = true
					break
				}
			}
		}
		if remainFuel != 0 && hasNext == false {
			if current != finish {
				mem[current][remainFuel] = 0
			} else {
				mem[current][remainFuel] = 1
			}
			return mem[current][remainFuel]
		}
		// 计算油量为 fuel，从位置 u 到 end 的路径数量
		// 由于每个点都可以经过多次，如果 u = end，那么本身就算一条路径
		var sum int
		if current == finish {
			sum++
		}
		for i := 0; i < len(locations); i++ {
			if i != current {
				need := cost(locations[current], locations[i])
				if remainFuel >= need {
					sum += traceback(i, finish, remainFuel-need)
					sum %= mod

				}
			}
		}
		mem[current][fuel] = sum
		return sum

	}
	return traceback(start, finish, fuel)
}

func cost(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func TestCountRoutes(t *testing.T) {
	//fmt.Println(countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5))
	fmt.Println(countRoutes2([]int{2, 3, 6, 8, 4}, 1, 3, 5))
	fmt.Println(countRoutes2([]int{1, 2, 3}, 0, 2, 40))
	//fmt.Println(countRoutes([]int{2, 1, 5}, 0, 0, 3))
	fmt.Println(countRoutes2([]int{2, 1, 5}, 0, 0, 3))

}
