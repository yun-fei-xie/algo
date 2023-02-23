package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"testing"
)

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
