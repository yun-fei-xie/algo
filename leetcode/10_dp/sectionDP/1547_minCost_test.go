package sectionDP

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/*
这个问题有比较多棘手的地方
1.
方法：
*/

/*
第一版写法：独立思考🤔
但是dfs函数四个参数，很难翻译成递推
*/
func minCost1(n int, cuts []int) int {
	sort.Ints(cuts)
	// left right : 木棍的左右端点
	// cLeft cRight: 当前需要考虑的切割点集合的左右区间端点
	var dfs func(left int, right int, cLeft int, cRight int) int
	dfs = func(left int, right int, cLeft int, cRight int) int {
		if cLeft > cRight {
			return 0
		}
		var c = math.MaxInt
		for i := cLeft; i <= cRight; i++ {
			cPos := cuts[i]
			m := (right - left) + dfs(left, cPos, cLeft, i-1) + dfs(cPos, right, i+1, cRight)
			c = min(c, m)
		}
		//fmt.Printf("left->%d  right->%d  cost->%d\n", left, right, c)
		return c
	}
	return dfs(0, n, 0, len(cuts)-1)
}

/*
在minCost1的基础上进行改进
将cuts数组排序后，在区间的首位分别添加上0和n，就变成了一根完整的木棍。
需要注意的是，(left,right) 开区间内的点才是切点。
在递归的过程中需要始终维持住这个定义。不然程序必然出错。
*/
func minCost2(n int, cuts []int) int {
	newCuts := make([]int, 0)
	sort.Ints(cuts)
	newCuts = append(newCuts, 0)
	newCuts = append(newCuts, cuts...)
	newCuts = append(newCuts, n)
	// [5,6,1,4,2]	-> [0,1,2,4,5,6,7] 一根完成的木棍 中间都是切割点
	// left right:newCuts数组下标
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left+1 == right {
			return 0
		}
		var m = math.MaxInt
		cost := newCuts[right] - newCuts[left]
		for i := left + 1; i < right; i++ {
			m = min(dfs(left, i)+dfs(i, right)+cost, m)
		}
		return m
	}
	return dfs(0, len(newCuts)-1)
}

/*
1:1动态规划
*/

func minCost3(n int, cuts []int) int {
	newCuts := make([]int, 0)
	sort.Ints(cuts)
	newCuts = append(newCuts, 0)
	newCuts = append(newCuts, cuts...)
	newCuts = append(newCuts, n)

	dp := make([][]int, len(newCuts))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(newCuts))
	}

	for left := len(newCuts) - 1; left >= 0; left-- {
		for right := left + 2; right < len(newCuts); right++ {
			var m = math.MaxInt
			var c = newCuts[right] - newCuts[left]
			for k := left + 1; k < right; k++ {
				m = min(m, c+dp[left][k]+dp[k][right])
			}
			dp[left][right] = m
		}
	}
	return dp[0][len(newCuts)-1]
}

func TestMinCost(t *testing.T) {
	fmt.Println(minCost1(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost2(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost3(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost1(7, []int{1, 3, 4, 5}))
	fmt.Println(minCost2(7, []int{1, 3, 4, 5}))
	fmt.Println(minCost3(7, []int{1, 3, 4, 5}))
}
