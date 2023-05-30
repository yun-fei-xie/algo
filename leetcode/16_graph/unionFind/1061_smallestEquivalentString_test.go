package unionFind_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
1061. 按字典序排列最小的等效字符串
https://leetcode.cn/problems/lexicographically-smallest-equivalent-string/?envType=study-plan-v2&envId=graph-theory

给出长度相同的两个字符串s1 和 s2 ，还有一个字符串 baseStr 。
其中  s1[i] 和 s2[i]  是一组等价字符。
举个例子，如果 s1 = "abc" 且 s2 = "cde"，那么就有 'a' == 'c', 'b' == 'd', 'c' == 'e'。
等价字符遵循任何等价关系的一般规则：
 自反性 ：'a' == 'a'
 对称性 ：'a' == 'b' 则必定有 'b' == 'a'
 传递性 ：'a' == 'b' 且 'b' == 'c' 就表明 'a' == 'c'
例如， s1 = "abc" 和 s2 = "cde" 的等价信息和之前的例子一样，那么 baseStr = "eed" , "acd" 或 "aab"，这三个字符串都是等价的，而 "aab" 是 baseStr 的按字典序最小的等价字符串
利用 s1 和 s2 的等价信息，找出并返回 baseStr 的按字典序排列最小的等价字符串。

方法：并查集
1.合并字符
2.遍历s1和s2收集同一个组中的字符
3.构造新的字符串
*/

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := InitUnionFind2(26)
	length := len(s1)
	for i := 0; i < length; i++ {
		uf.unionElements(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	// 收集
	charSet := make(map[int][]uint8, 0)
	for i := 0; i < length; i++ {
		root := uf.find(int(s1[i] - 'a'))
		charSet[root] = append(charSet[root], s1[i])
		charSet[root] = append(charSet[root], s2[i])
	}

	// 去重
	for root, chars := range charSet {
		set := make(map[uint8]struct{})
		for _, char := range chars {
			set[char] = struct{}{}
		}
		cc := make([]uint8, 0)
		for char, _ := range set {
			cc = append(cc, char)
		}
		sort.Slice(cc, func(i, j int) bool {
			return cc[i] < cc[j]
		})
		charSet[root] = cc
	}

	// 找等价
	baseStrLength := len(baseStr)
	sb := strings.Builder{}
	for i := 0; i < baseStrLength; i++ {
		root := uf.find(int(baseStr[i] - 'a'))

		if _, found := charSet[root]; found {
			sb.WriteByte(charSet[root][0])
		} else {
			sb.WriteByte(baseStr[i])
		}
	}
	return sb.String()
}

func TestSmallestEquivalentString(t *testing.T) {
	fmt.Println(smallestEquivalentString("hello", "world", "hold"))
	fmt.Println(smallestEquivalentString("leetcode", "programs", "sourcecode"))
}
