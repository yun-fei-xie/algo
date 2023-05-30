package unionFind_test

import (
	"fmt"
	"testing"
)

/*
990. 等式方程的可满足性
https://leetcode.cn/problems/satisfiability-of-equality-equations/description/

给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

方法：并查集
1.等式的两端的元素一定要在同一个连通分量中。
2.不等式的两端的元素不能在同一个连通分量中。
3.遍历方程，对等式两端的元素进行union操作。完成并查集的并
4.遍历方程，对于所有的等式，两端元素一定要在同一个集合中，对于不等式，两端的元素一定不能在同一个集合中。
*/
func equationsPossible(equations []string) bool {
	uf := InitUnionFind2(26)
	for _, equation := range equations {
		if equation[1:len(equation)-1] == "==" {
			v1 := int(equation[0] - 'a')
			v2 := int(equation[len(equation)-1] - 'a')
			uf.unionElements(v1, v2)
		}
	}

	for _, equation := range equations {
		v1 := int(equation[0] - 'a')
		v2 := int(equation[len(equation)-1] - 'a')

		v1Root := uf.find(v1)
		v2Root := uf.find(v2)

		if equation[1:len(equation)-1] == "==" {

			if v1Root != v2Root {
				return false
			}
		} else if equation[1:len(equation)-1] == "!=" {
			if v1Root == v2Root {
				return false
			}
		}
	}
	return true
}

func TestEquationsPossible(t *testing.T) {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
}
