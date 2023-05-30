package _6_graph

import (
	"fmt"
	"testing"
)

/*


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
