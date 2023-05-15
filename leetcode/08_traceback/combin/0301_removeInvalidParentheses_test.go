package combin

import (
	"fmt"
	"testing"
)

/*
301.删除无效的括号
https://leetcode.cn/problems/remove-invalid-parentheses/

给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
返回所有可能的结果。答案可以按 任意顺序 返回。

示例 1：

输入：s = "()())()"
输出：["(())()","()()()"]
示例 2：

输入：s = "(a)())()"
输出：["(a())()","(a)()()"]
示例 3：

输入：s = ")("
输出：[""]

方法：
考虑枚举s中的每一个字符。
对于s[i]这个字符来说，它要么被加入结果集合中，要么被抛弃（删除掉）。
 1. s[i]是非括号字符。->直接添加到结果集合中。
 2. s[i]是左括号->它能不能被放入到结果集合中呢？ 需要看结果集合中最多可以放入多少个左括号。（这是一个考点 这个地方容易出错）
    然后利用22题括号生成中的性质进行判断。
 3. s[i]是有括号->它能不能放入到结果集合中呢？需要看结果集合中当前有括号的数量是不是<左括号。（如果是，就可以尝试放入）
 4. 步骤2和步骤3其实是在回溯。如果用树形来看的话，是同一个节点的两个孩子节点。
*/
func removeInvalidParentheses(s string) []string {

	leftParentheses, rightParentheses := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftParentheses++
		} else if s[i] == ')' {
			rightParentheses++
		}
	}

	maxParentheses := min(leftParentheses, rightParentheses)
	var path = make([]uint8, 0)
	var res = make(map[string]struct{})
	var ans = make([]string, 0)
	var maxLen = 0
	var dfs func(index int, left int, right int)
	dfs = func(index int, left int, right int) {
		if right > left || left > maxParentheses {
			return
		}
		if index == len(s) {
			if len(path) == 0 || left != right {
				return
			} else if len(path) > maxLen {
				maxLen = len(path)
				temp := make([]uint8, len(path))
				copy(temp, path)
				res = map[string]struct{}{string(temp): {}}
			} else if len(path) == maxLen {
				temp := make([]uint8, len(path))
				copy(temp, path)
				res[string(temp)] = struct{}{}
			}
			return
		}

		ch := s[index]
		if ch != '(' && ch != ')' {
			path = append(path, ch)
			dfs(index+1, left, right)
			path = path[0 : len(path)-1]
			// 这里必须回溯 例如：["(f"]不合法，从下层回来，应该将f去掉，然后回到它的上一层["("]
			// 上一走尝试不加"("的路线，于是递归到i指向f的位置。
		}

		if ch == '(' {
			path = append(path, '(')
			dfs(index+1, left+1, right)
			path = path[0 : len(path)-1]
			dfs(index+1, left, right)
		}

		if ch == ')' {
			path = append(path, ')')
			dfs(index+1, left, right+1)
			path = path[0 : len(path)-1]
			dfs(index+1, left, right)
		}
	}
	dfs(0, 0, 0)
	if maxLen == 0 {
		return []string{""}
	}
	for k, _ := range res {
		ans = append(ans, k)
	}
	return ans
}

func min(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

func TestRemoveInvalidParentheses(t *testing.T) {
	//fmt.Println(removeInvalidParentheses("()())()"))
	//fmt.Println(removeInvalidParentheses("(a)())()"))
	//fmt.Println(removeInvalidParentheses(")("))
	//fmt.Println(removeInvalidParentheses("x("))
	fmt.Println(removeInvalidParentheses(")(f"))
}
