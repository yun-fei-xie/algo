package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
just two chars
use stack
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

这一个极简消消乐模型,体会一下。
当两个字符匹配到了就进行消除操作。 匹配的定义就是（两个字符相同）
整个匹配消除使用栈来实现。
*/
func removeDuplicates(s string) string {

	stack := list.New()
	sb := strings.Builder{}
	// 匹配消除阶段
	for i := 0; i < len(s); i++ {
		char := s[i]

		if stack.Len() == 0 {
			stack.PushBack(char)
			continue
		} else {
			element := stack.Back()
			if element.Value.(uint8) == char {
				stack.Remove(element)
			} else {
				stack.PushBack(char)
			}
		}
	}
	// 收集栈中将没有被消除的字符，将其放入到结果中。
	for stack.Len() != 0 {
		front := stack.Front()
		sb.WriteString(string(front.Value.(uint8)))
		stack.Remove(front)
	}
	return sb.String()
}

func TestRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	res := removeDuplicates(s)
	fmt.Println(res)
}
