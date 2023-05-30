package unionFind_test

import (
	"fmt"
	"testing"
)

/*
方法：求解图的连通分量的个数。如果两个字符串相似，那么两个字符串可以连接成一条边。

	如何判断相似？如果s1和s2只有两个字符不同s1[i]!=s2[i] s1[j]!=s2[j] 同时，s1[i]==s2[j] s2[i]==s1[j]
	这里不需要连边，直接上并查集。
*/
func numSimilarGroups1(strs []string) int {
	nodeCount := len(strs)
	uf := InitUnionFind2(nodeCount)
	for i := 0; i < nodeCount; i++ {
		for j := i + 1; j < nodeCount; j++ {
			if isSimilar(strs[i], strs[j]) {
				uf.unionElements(i, j)
			}
		}
	}

	ans := make(map[int]struct{})
	for i := 0; i < nodeCount; i++ {
		root := uf.find(i)
		ans[root] = struct{}{}
	}
	return len(ans)
}

func isSimilar(s1 string, s2 string) bool {
	count := 0
	length := len(s1)
	first, second := -1, -1
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			if first == -1 {
				first = i
				count++
			} else {
				second = i
				count++
			}
		}
	}
	if count == 0 {
		return true
	}
	if count == 2 {
		if s1[first] == s2[second] && s1[second] == s2[first] {
			return true
		}
	}
	return false
}

func TestIsSimilar(t *testing.T) {
	fmt.Println(numSimilarGroups1([]string{"tars", "rats", "arts", "star"}))
	fmt.Println(numSimilarGroups1([]string{"omv", "ovm"}))
}
