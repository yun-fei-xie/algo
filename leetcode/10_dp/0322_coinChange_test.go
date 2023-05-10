package _0_dp

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/coin-change/?favorite=2cktkvj

零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

思路：
1.贪心算法 先兑换面值大的 (贪心不行) 因为兑换完大面额的，可会会出现无解
2.暴力枚举，当然，每一个面值的枚举的数量都有一个上限。这个上限就是当前余额amount 除以 当前硬币的面值，就得到一个最大的枚举数量
暴力枚举会超时。 是否可以进行优化->记忆化搜索
3.记忆化搜索？



*/

func coinChange2(coins []int, amount int) int {
	var res = math.MaxInt64
	var dfs func(coinIndex int, remainValue int, coinAmount int)
	dfs = func(coinIndex int, remainValue int, coinAmount int) {
		if coinIndex >= len(coins) || remainValue == 0 {
			if remainValue == 0 && coinAmount < res { // 找到了一个解
				res = coinAmount
			}
			return
		}

		// 一枚一枚地尝试
		// i -> [0...remainValue/coinValue]  i表示本次使用的硬币的枚数
		for i := 0; i <= remainValue/coins[coinIndex]; i++ {
			dfs(coinIndex+1, remainValue-i*coins[coinIndex], coinAmount+i)
		}

	}

	dfs(0, amount, 0)
	if res != math.MaxInt64 {
		return res
	}
	return -1
}

/*
错误解法: 这个题不能用贪心算法
*/
func coinChange(coins []int, amount int) int {
	res := 0

	sort.Slice(coins, func(i, j int) bool {
		if coins[i] > coins[j] {
			return true
		}
		return false
	})

	for i := 0; i < len(coins); {
		value := coins[i]
		if value > amount { // 当前这个面值大了，需要换个小点的进行兑换
			i++
			continue
		} else {
			cnt := amount / value
			res += cnt
			amount -= cnt * value
			if amount == 0 {
				break
			}
			i++
		}
	}

	if amount != 0 { // 兑换不完
		return -1
	} else {
		return res
	}

}

func TestCoinChange(t *testing.T) {
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	//fmt.Println(coinChange([]int{2}, 3))
	//fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange2([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange2([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
}
