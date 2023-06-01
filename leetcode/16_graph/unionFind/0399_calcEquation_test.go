package unionFind

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

方法：带权并查集(第一次碰到带权并查集)
1. 在常规的并查集中加入边权
2. equations中的元素可能会带有多个字母，但是题目给了(每个 Ai 或 Bi 是一个表示单个变量的字符串)，这个时候不需要做字符约分之类的操作，直接看做整体变量即可。
*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	length := len(equations)
	uf := newUf(length * 2)
	// 编码 合并
	id := 0
	hashMap := make(map[string]int)
	for i := 0; i < length; i++ {
		p := equations[i][0]
		q := equations[i][1]
		w := values[i]
		if _, found := hashMap[p]; !found {
			hashMap[p] = id
			id++
		}
		if _, found := hashMap[q]; !found {
			hashMap[q] = id
			id++
		}
		uf.unionElements(hashMap[p], hashMap[q], w)
	}

	// 查询
	ans := make([]float64, 0)
	for _, query := range queries {
		p := query[0]
		q := query[1]

		pid, foundP := hashMap[p]
		qid, foundQ := hashMap[q]

		if !foundP || !foundQ {
			ans = append(ans, -1.0)
		} else {
			w := uf.isConnected(pid, qid)
			ans = append(ans, w)
		}
	}
	return ans
}

// 带权并查集
// 如何对equation进行编码？按照方程的下标进行编码。
// 使用hash表和自增ID对每一个元素进行编码。这样不需要考虑元素重复的问题。
type ufWeighted struct {
	parent []int
	weight []float64
}

func newUf(n int) *ufWeighted {
	parent := make([]int, n)
	weight := make([]float64, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		weight[i] = 1.0
	}
	return &ufWeighted{
		parent: parent,
		weight: weight,
	}
}

/*
路径压缩查找
*/
func (uf *ufWeighted) find(p int) int {
	if uf.parent[p] != p {
		par := uf.parent[p]
		uf.parent[p] = uf.find(par)
		// 更新权重
		uf.weight[p] *= uf.weight[par]
	}
	return uf.parent[p]
}

/*
计算p/q这个表达式的值
执行find会自动进行路径压缩，使得p和q直接和根节点进行相连
*/
func (uf *ufWeighted) isConnected(p int, q int) float64 {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return uf.weight[q] / uf.weight[p]
	}
	return -1.0
}

/*
让被除数作为根节点
*/
func (uf *ufWeighted) unionElements(p int, q int, weight float64) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	uf.parent[rootQ] = rootP
	uf.weight[rootQ] = uf.weight[p] * weight / uf.weight[q]

}

func TestCalcEquation(t *testing.T) {
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}))
}
