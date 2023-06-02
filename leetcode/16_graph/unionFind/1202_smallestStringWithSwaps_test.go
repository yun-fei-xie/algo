package unionFind_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
1202. 交换字符串中的元素
https://leetcode.cn/problems/smallest-string-with-swaps/description/

给你一个字符串 s，以及该字符串中的一些「索引对」数组 pairs，其中 pairs[i] = [a, b] 表示字符串中的两个索引（编号从 0 开始）。
你可以 任意多次交换 在 pairs 中任意一对索引处的字符。
返回在经过若干次交换后，s 可以变成的按字典序最小的字符串。

方法：并查集
1. 对pairs中给出的位置进行合并，会得到一个或者多个组。
2. 遍历下标，将同一组的字符收集并排序。
3. 遍历下标，生成答案。对于下标i，找出root并将组内的第一个元素放在该位置上，然后将这个刚刚放置的字符从组内删除。（删除就不用维护下标）
*/
func smallestStringWithSwaps(s string, pairs [][]int) string {

	uf := newSmaUf(len(s))

	for _, pair := range pairs {
		uf.unionElements(pair[0], pair[1])
	}
	group := make(map[int][]uint8)
	for i := 0; i < len(s); i++ {
		root := uf.find(i)
		group[root] = append(group[root], s[i])
	}

	for _, v := range group {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	var ans = make([]uint8, len(s))

	for i := 0; i < len(s); i++ {
		root := uf.find(i)
		ans[i] = group[root][0]
		group[root] = group[root][1:]
	}
	return string(ans)
}

type smaUf struct {
	parent []int
}

func newSmaUf(n int) *smaUf {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &smaUf{parent: parent}
}

func (uf *smaUf) find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *smaUf) unionElements(p int, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
}

func TestSmallestStringWithSwap(t *testing.T) {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
	fmt.Println(smallestStringWithSwaps("udyyek", [][]int{{3, 3}, {3, 0}, {5, 1}, {3, 1}, {3, 4}, {3, 5}}))
}
