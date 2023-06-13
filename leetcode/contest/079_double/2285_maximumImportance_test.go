package _79_double

import "sort"

/*
2285. 道路的最大总重要性
https://leetcode.cn/problems/maximum-total-importance-of-roads/description/

方法：出度越大，权重越大。用出度分配权重
*/
func maximumImportance(n int, roads [][]int) int64 {
	outDegree := make([]int, n)
	for _, road := range roads {
		c1 := road[0]
		c2 := road[1]
		outDegree[c1]++
		outDegree[c2]++
	}

	sort.Ints(outDegree)
	var ans int
	for i := 1; i <= n; i++ {
		ans += i * outDegree[i-1]
	}
	return int64(ans)
}
