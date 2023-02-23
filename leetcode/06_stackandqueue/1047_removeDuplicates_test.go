package _6_stackandqueue

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

/*
just two chars
use stack
*/
func removeDuplicates(s string) string {

	stack := list.New()
	sb := strings.Builder{}
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
