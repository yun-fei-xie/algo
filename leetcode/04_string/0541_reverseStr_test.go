package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
*
https://leetcode.cn/problems/reverse-string-ii/description/
开始还没什么思路 反正就是模拟
网站上的代码比我的要简洁
https://programmercarl.com/0541.%E5%8F%8D%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2II.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
*/
func reverseStr(s string, k int) string {
	reverse := true // 翻转->不翻转->翻转...
	sb := strings.Builder{}
	var i int = 0
	for i = 0; i < len(s); i = i + k {
		if reverse {
			reverseSubString := reverseS(s[i:min(i+k, len(s))]) // [left , right)
			sb.WriteString(reverseSubString)
			reverse = false
		} else {
			sb.WriteString(s[i:min(i+k, len(s))])
			reverse = true
		}
	}

	//sb.WriteString(s[i-k : min(i, len(s))])

	return sb.String()
}

// 这里把参数修改为[]byte 就可以使用双指针的方式翻转传入的字符串
func reverseS(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}

}

func TestReverseStr(t *testing.T) {
	s := "abcdefg"
	k := 2

	res := reverseStr(s, k)
	fmt.Println(res)
}

func TestStringCopy(t *testing.T) {

	s := "abcde"
	s1 := s[0:len(s)]
	fmt.Println(s1)

}
