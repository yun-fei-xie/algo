package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode.cn/problems/reverse-words-in-a-string/description/
这个题目比较麻烦的是，在s中，单词与单词之间的空格可能会有多个。
调用了库函数
对于strings的切割函数 返回值有点奇怪
*/
func reverseWords(s string) string {
	sb := strings.Builder{}
	splitStrings := strings.Fields(s)
	for i := len(splitStrings) - 1; i >= 0; i-- {
		sb.WriteString(splitStrings[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func TestReverseWords(t *testing.T) {
	s := "a good   example"
	res := reverseWords(s)
	fmt.Println(res)

	//testString := "  aa   bbb   ccc   "
	//res := strings.Fields(testString)
	//for _, s := range res {
	//	fmt.Print(s)
	//}

}
