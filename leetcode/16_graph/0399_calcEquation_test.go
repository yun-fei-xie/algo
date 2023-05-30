package _6_graph

import (
	"fmt"
	"testing"
)

/*
399. 除法求值
https://leetcode.cn/problems/evaluate-division/

给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

方法：图问题
1.条件里面的顶点用的是字符串表示，不太容易直接编号。考虑用map[string][]string
2.由于是带权图，需要存储权重，考虑使用map[string]map[string]float64 第一个string表示当前顶点，第二个string表示和这个顶点相连接的顶点，float64表示两个顶点之间的边的权值

*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	graph := make(map[string]map[string]float64, 0)
	visited := make(map[string]bool)
	for i := 0; i < len(equations); i++ {
		v1 := equations[i][0]
		v2 := equations[i][1]
		w := values[i]

		if _, found := graph[v1]; !found {
			graph[v1] = make(map[string]float64)
		}
		if _, found := graph[v2]; !found {
			graph[v2] = make(map[string]float64)
		}
		graph[v1][v2] = w
		graph[v2][v1] = 1 / w

		visited[v1] = false
		visited[v2] = false
	}
	// 起点->终点 路径权值的乘积
	var ans = make([]float64, 0)
	for _, query := range queries {
		q1 := query[0]
		q2 := query[1]
		_, found1 := graph[q1]
		_, found2 := graph[q2]
		if !found1 || !found2 {
			ans = append(ans, -1.0)
			continue
		}
		if q1 == q2 {
			ans = append(ans, 1.0)
			continue
		}

		// 搜索路径
		ans = append(ans, dfs(graph, q1, q2))
	}
	return ans
}

func dfs(graph map[string]map[string]float64, start string, end string) (ans float64) {
	return -1
}

func TestCalcEquation(t *testing.T) {
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
}
