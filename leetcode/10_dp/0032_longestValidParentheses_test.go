package _0_dp

import (
	"container/list"
	"fmt"
	"testing"
)

/*
32. 最长有效括号
https://leetcode.cn/problems/longest-valid-parentheses/description/
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

注意点：本题求解的是最长合法子串，而不是子序列。
方法：栈
*/

/*
用栈模拟一遍
栈里面得存储下标
*/
func longestValidParentheses(s string) int {
	arr := make([]int, len(s))
	stack := list.New()
	ans := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '(' {
			stack.PushBack(i)
		} else if ch == ')' {
			if stack.Len() == 0 {
				arr[i] = 1
			} else {
				stack.Remove(stack.Back())
			}
		}
	}
	for stack.Len() != 0 {
		arr[stack.Back().Value.(int)] = 1
		stack.Remove(stack.Back())
	}

	zeroCnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			ans = max(ans, zeroCnt)
			zeroCnt = 0
		} else {
			zeroCnt++
		}
	}
	ans = max(ans, zeroCnt)
	return ans
}

/*
dp[0...i]以i为结尾字符的最长合法子串
以后再做
*/
func longestValidParentheses2(s string) int {
	return -1
}

func TestLongestValidParentheses(t *testing.T) {

	//fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("(()"))
}
