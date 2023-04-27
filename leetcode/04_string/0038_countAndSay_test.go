package _4_string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
初始时从'1'开始

1112234445666
使用双指针来分段统计连续数字字符出现的次数。
*/
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		sb := strings.Builder{}
		for start, cursor := 0, 0; cursor < len(prev); start = cursor {
			for cursor < len(prev) && prev[cursor] == prev[start] {
				cursor++
			}
			sb.WriteString(strconv.Itoa(cursor - start))
			sb.WriteByte(prev[start])
		}
		prev = sb.String()
	}
	return prev
}

func TestCountAndSay(t *testing.T) {
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}
