package _0_dp

import (
	"fmt"
	"testing"
)

/*
2140.解决智力问题
https://leetcode.cn/problems/solving-questions-with-brainpower/description/?envType=study-plan-v2&envId=dynamic-programming

给你一个下标从 0 开始的二维整数数组 questions ，其中 questions[i] = [pointsi, brainpoweri] 。
这个数组表示一场考试里的一系列题目，你需要 按顺序 （也就是从问题 0 开始依次解决），针对每个问题选择 解决 或者 跳过 操作。解决问题 i 将让你 获得  pointsi 的分数，但是你将 无法 解决接下来的 brainpoweri 个问题（即只能跳过接下来的 brainpoweri 个问题）。如果你跳过问题 i ，你可以对下一个问题决定使用哪种操作。

比方说，给你 questions = [[3, 2], [4, 3], [4, 4], [2, 5]] ：
如果问题 0 被解决了， 那么你可以获得 3 分，但你不能解决问题 1 和 2 。
如果你跳过问题 0 ，且解决问题 1 ，你将获得 4 分但是不能解决问题 2 和 3 。
请你返回这场考试里你能获得的 最高 分数。

这个题很像买卖股票包含冷冻期哪个题目(309号题目)
*/
func mostPoints1(questions [][]int) int64 {
	var ans int64
	var dfs func(i int, cool int, point int64)
	dfs = func(i int, cool int, point int64) {
		if i == len(questions) {
			if point > ans {
				ans = point
			}
			return
		}
		//如果当前能选择
		if cool == 0 {
			// 选择
			dfs(i+1, cool+questions[i][1], point+int64(questions[i][0]))
			// 不选择
			dfs(i+1, cool, point)
		} else {
			// 当前不能选择
			dfs(i+1, cool-1, point)
		}
	}
	dfs(0, 0, 0)
	return ans
}

/*
这样思考🤔，当前的选择会影响后面的选择。
那么枚举当前的选择
这个写法还有优化的空间。因为cool的值是知道的，那么可以直接跳过后续的部分问题
*/
func mostPoints2(questions [][]int) int64 {
	// i是当前需要考虑问题，cool表示当前是否需要冷却，0不需要
	// 返回questions[i...len-1]这个区间能够获得的点数的最大值
	var dfs func(i int, cool int) int64
	dfs = func(i int, cool int) int64 {
		if i == len(questions) {
			return 0
		}
		if cool > 0 {
			// 不能选
			return dfs(i+1, cool-1)
		} else {
			// 可以选（包含两种情况：选和不选）
			return max64(dfs(i+1, 0), dfs(i+1, questions[i][1])+int64(questions[i][0]))
		}
	}
	return dfs(0, 0)
}

func max64(args ...int64) int64 {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

/*
翻译成递推
这个题翻译成递推
https://leetcode.cn/problems/solving-questions-with-brainpower/solutions/1213919/dao-xu-dp-by-endlesscheng-2qkc/
根据数据条件，最长冷却100000
*/
func mostPoints4(questions [][]int) int64 {
	dp := make([]int64, len(questions)+1)
	for i := len(questions) - 1; i >= 0; i-- {
		q := questions[i]
		// 选
		// 如果选，那么下一个可以被解决的问题是 i + q[1] + 1
		if j := i + q[1] + 1; j < len(questions) {
			dp[i] = max64(dp[i+1], dp[j]+int64(q[0]))
		} else {
			// 这个地方包含了一种情况:j越界
			dp[i] = max64(dp[i+1], int64(q[0]))
		}
	}
	return dp[0]
}

func TestMostPoints(t *testing.T) {
	fmt.Println(mostPoints1([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	fmt.Println(mostPoints2([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
}
