package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
如何识别宫格字符
*/

func replaceSpace(s string) string {
	b := []byte(s)
	sb := strings.Builder{}
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' { // 识别空格字符
			sb.WriteString("%20")
		} else {
			sb.WriteByte(b[i])
		}
	}
	return sb.String()
}

// 使用库函数
func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	res := replaceSpace(s)
	fmt.Println(res)
}
