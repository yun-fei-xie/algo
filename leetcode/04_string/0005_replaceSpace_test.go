package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*

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

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	res := replaceSpace(s)
	fmt.Println(res)
}
