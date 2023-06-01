package _25

import (
	"fmt"
	"math"
	"testing"
)

/*
方法：暴力求解
1.计算组合
2.计算一个组合中任意两数之差的最小值
3.求出所有组合中的最小值中的最大值
时间和空间都不符合要求

方法：二分
https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/solutions/2031994/er-fen-da-an-by-endlesscheng-r418/




*/

func maximumTastiness(price []int, k int) int {
	combine := make([][]int, 0)
	path := make([]int, 0)
	var traceback func(i int, depth int)
	traceback = func(i int, depth int) {
		if depth == k {
			dist := make([]int, len(path))
			copy(dist, path)
			combine = append(combine, dist)
			return
		}

		for j := i; j < len(price); j++ {
			if len(price)-j+1+depth < k {
				continue
			}
			path = append(path, price[j])
			traceback(j+1, depth+1)
			path = path[:len(path)-1]
		}
	}
	traceback(0, 0)

	return calc(combine)
}

func calc(baskets [][]int) int {
	maxTastiness := math.MinInt
	for _, basket := range baskets {
		minDiff := math.MaxInt
		for i := 0; i < len(basket); i++ {
			for j := i + 1; j < len(basket); j++ {
				minDiff = min(minDiff, diffAbs(basket[i], basket[j]))
			}
		}
		maxTastiness = max(maxTastiness, minDiff)
	}
	return maxTastiness
}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}
func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func diffAbs(i int, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
func TestMaximumTastiness(t *testing.T) {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)) // c6-3 -> (6*5*4)/3*2 = 20
	fmt.Println(maximumTastiness([]int{1, 3, 1}, 2))
	fmt.Println(maximumTastiness([]int{7, 7, 7, 7}, 2))
}
