package sectionDP

import (
	"fmt"
	"math"
	"testing"
)

/*
这个题目有一个处理难点在于，气球打破之后，相邻的气球在nums数组中并不相邻。
*/
func maxCoins1(nums []int) int {
	var flag = make([]bool, len(nums))
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left > right {
			return 0
		}
		var m = math.MinInt
		for i := left; i <= right; i++ {
			leftCoin, rightCoin := 1, 1
			for j := i - 1; j >= left; j-- {
				if flag[j] == false {
					leftCoin = nums[j]
					break
				}
			}

			for j := i + 1; j <= right; j++ {
				if flag[j] == false {
					rightCoin = nums[j]
					break
				}
			}
			flag[i] = true

			p0 := leftCoin * nums[i] * rightCoin
			p1 := dfs(left, i-1)
			p2 := dfs(i+1, right)
			coin := p0 + p1 + p2

			//coin := leftCoin*nums[i]*rightCoin + dfs(left, i-1) + dfs(i+1, right)
			flag[i] = false
			m = max(m, coin)
		}
		return m
	}
	return dfs(0, len(nums)-1)
}

/*
重新组织思路
这个代码写的dp不了
*/
func maxCoins2(nums []int) int {
	var flag = make([]bool, len(nums))

	var dfs func(d int) int
	dfs = func(d int) int {
		if d >= len(nums) {
			return 0
		}
		var m = math.MinInt
		for i := 0; i < len(flag); i++ {
			if flag[i] == false {
				// 戳破
				flag[i] = true
				leftCoins, rightCoins := 1, 1
				for j := i - 1; j >= 0; j-- {
					if flag[j] == false {
						leftCoins = nums[j]
						break
					}
				}
				for j := i + 1; j < len(nums); j++ {
					if flag[j] == false {
						rightCoins = nums[j]
						break
					}
				}

				coins := nums[i]*leftCoins*rightCoins + dfs(d+1)
				m = max(m, coins)
				flag[i] = false
			}
		}
		return m
	}
	return dfs(0)
}

func TestMaxCoins(t *testing.T) {
	fmt.Println(maxCoins2([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins2([]int{1, 5}))
}
