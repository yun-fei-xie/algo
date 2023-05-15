package combin

import (
	"fmt"
	"testing"
)

/*
22. 括号生成
https://leetcode.cn/problems/generate-parentheses/

["((()))","(()())","(())()","()(())","()()()"]

1. 最终的效果是：左括号的数量=有括号的数量
2. 在生成的过程中：左括号的数量>=有括号的数量
这个规律可以通过自己手写括号得出。

观察合法的括号，可以发现，如果括号的总的字符数量是2n,在括号生成的过程中:
1. 对于每一个合法的括号字符串的前缀，左括号的个数>=右括号的个数
2. 右括号的个数必须小于或者等于左括号的个数。如果当前左右括号相等，就只能选择左括号。
*/
func generateParenthesis(n int) []string {

	var path = make([]uint8, 2*n)
	var ans = make([]string, 0)
	var m = 2 * n
	var dfs func(i int, left int)
	dfs = func(i int, left int) {
		if i == m {
			temp := make([]uint8, len(path))
			copy(temp, path)
			ans = append(ans, string(temp))
		}
		// 左括号的数量<n
		if left < n {
			// 尝试左括号
			path[i] = '('
			dfs(i+1, left+1)
		}
		// 这个地方在回溯 两个if 会同时成立
		// 所以，这就是在枚举这个位置上可能出现的情况。
		if i-left < left {
			path[i] = ')'
			dfs(i+1, left)
		}
	}
	dfs(0, 0)
	return ans
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
