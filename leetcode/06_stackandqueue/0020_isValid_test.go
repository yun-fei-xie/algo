package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/valid-parentheses/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


解题思路：使用栈进行匹配。如果是左括号，则将当前元素压入栈中。
	如果是右括号，则将栈顶元素弹出尝试进行匹配。

*/

func isValid(s string) bool {

	stack := list.New()

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' || char == '{' || char == '[' {
			stack.PushBack(char)
		} else {

			if stack.Len() == 0 {
				return false
			}
			element := stack.Back()
			left := element.Value.(uint8)

			if left == '(' && char != ')' || left == '[' && char != ']' || left == '{' && char != '}' {
				return false
			}
			stack.Remove(element)
		}
	}
	return stack.Len() == 0
}

func TestIsValid(t *testing.T) {
	s := "()[]{}"
	s2 := "([)]"
	res := isValid(s)
	res2 := isValid(s2)
	fmt.Println(res)
	fmt.Println(res2)
}
