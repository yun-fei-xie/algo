package sectionDP

import (
	"fmt"
	"math"
	"testing"
)

/*
1039. 多边形三角剖分的最低得分
https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/

你有一个凸的 n 边形，其每个顶点都有一个整数值。给定一个整数数组 values ，其中 values[i] 是第 i 个顶点的值（即 顺时针顺序 ）。
假设将多边形 剖分 为 n - 2 个三角形。对于每个三角形，该三角形的值是顶点标记的乘积，三角剖分的分数是进行三角剖分后所有 n - 2 个三角形的值之和。
返回 多边形进行三角剖分后可以得到的最低分 。

示例 1：
输入：values = [1,2,3]
输出：6
解释：多边形已经三角化，唯一三角形的分数为 6。

方法：递归or区间dp
	递归或者动态规划的难点都是如何正确的切分子问题。
	这样思考，初始时，从区间的左端点i,到区间的右端点j表示一个多边形。它可以认为是顺指针从i走到j i->j
	然后再从j直接走到i，j->i。（j->i是一条直接的边)
	考虑从i到j之间选择一个顶点k,(i<k<j)，那么（i,k,j）三个点就可以形成一个三角形。同时还会将原来的多边形
	切分成多边形（j...k）和多边形（k...i）和三角形（i,k,j）。也就是选择一个点之后，会得到两个多边形和一个三角形。
	对于这个三角形可以直接求解，对于新得到的两个多边形可以进行递归求解。于是，k是选出来的一个值，因此可以在限定的区间中枚举k。

	递归出口：当i+1==j的时候，可以发现，区间中没有k可以进行枚举，无法构成三角形，递归到底。
*/

func minScoreTriangulation1(values []int) int {

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i+1 == j {
			return 0
		}

		var m = math.MaxInt
		for k := i + 1; k < j; k++ {
			m = min(m, dfs(i, k)+dfs(k, j)+values[i]*values[k]*values[j])
		}
		return m
	}
	return dfs(0, len(values)-1)
}

/*
将递归翻译成递推
本地的递推方向容易出错。
*/
func minScoreTriangulation2(values []int) int {
	length := len(values)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	// 递推方向：初始时，i,j为问题的最小规模（length-3,i+2）这其实是一个三角形
	for i := length - 3; i >= 0; i-- {
		for j := i + 2; j < length; j++ {

			var m = math.MaxInt
			for k := i + 1; k < j; k++ {
				m = min(m, dp[i][k]+dp[k][j]+values[i]*values[k]*values[j])
			}
			dp[i][j] = m
		}
	}
	return dp[0][length-1]
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

func TestMinScoreTriangulation(t *testing.T) {
	fmt.Println(minScoreTriangulation1([]int{3, 7, 4, 5}))
	fmt.Println(minScoreTriangulation2([]int{3, 7, 4, 5}))
}
