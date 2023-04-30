package _0_dp

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
https://leetcode.cn/problems/russian-doll-envelopes/
*/
/*
暴力回溯求解可以通过69/87
*/
func maxEnvelopes(envelopes [][]int) int {
	var num = len(envelopes)
	var max = math.MinInt32
	var used = make([]bool, num)
	// dfs 以位置i为最小信封 搜索套娃信封的数量 递归开始时，current位置的信封已经放到了套娃信封里
	var dfs func(current int, sum int)
	dfs = func(current int, sum int) {
		//遍历其他还没有被使用过的信封 尝试选一个信封放进去
		var hasNext bool
		for i := 0; i < num; i++ {
			if used[i] == false && canContain(envelopes[current][0], envelopes[current][1], envelopes[i][0], envelopes[i][1]) {
				dfs(i, sum+1)
				hasNext = true
			}
		}
		// 如果一个信封都放不进去，这条线路终止
		if hasNext == false {
			if sum > max {
				max = sum
			}
		}
	}
	for i := 0; i < num; i++ {
		dfs(i, 1)
	}
	return max
}

// canContain 规格为w1、h1的信封是否可以容纳规格为w2、h2的信封
func canContain(w1, h1, w2, h2 int) bool {
	if w1 > w2 && h1 > h2 {
		return true
	}
	return false
}

/*
这个问题的本质就是一个最长上升子序列问题。
把二维信封看做是一个元素，就可以将这个问题抽象到leetcode第300号问题
如何定义大于，如果信封A可以包含另外一个信封B，那么信封A就大于信封B。

	[[5,4],[6,4],[6,7],[2,3]]

dp   2      2       2     1

但是这样做发现不对。例如，[5,4]可以放在[6,7]和[2,3]之间，构成[6,7]->[5,4]->[2,3]序列。
最长递增子序列的要求是子序列（子序列中的元素顺序相对于原始序列不的相对位置不能发生改变。）
但是对于套娃来说，没有这个限制。套娃不能收到相对位置这个限制。

达到长度为3的最长套娃。那我考虑除了当前元素外的所有元素，用一个环形来记性处理不行吗？
确实不行。环形同样有顺序性。在进行考虑的时候，被元素的相对位置影响

	[[30,50],[12,2],[3,4],[12,15]]

->  [[30,50],[12,15],[12,2],[3,4]]
dp      3         2       2     1

动态规划后，85/87 算是通过吧
*/
func maxEnvelopes2(envelopes [][]int) int {
	num := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1] // 如果元素i的第0个元素等于元素j的第0个元素，则第二个元素大的排在前面
		} else {
			return envelopes[i][0] > envelopes[j][0] // 否则的话，第0个元素大的排在前面
		}
	})
	var dp = make([]int, num)
	for i := 0; i < num; i++ {
		dp[i] = 1
	}

	var ans = math.MinInt32

	for i := num - 1; i >= 0; i-- {
		for j := i + 1; j < num; j++ {
			if canContain(envelopes[i][0], envelopes[i][1], envelopes[j][0], envelopes[j][1]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}

func TestMaxEnvelopes(t *testing.T) {
	//fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	//fmt.Println(maxEnvelopes2([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	//fmt.Println(maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
	//fmt.Println(maxEnvelopes2([][]int{{1, 1}, {1, 1}, {1, 1}}))
	fmt.Println(maxEnvelopes2([][]int{{30, 50}, {12, 2}, {3, 4}, {12, 15}}))
}
