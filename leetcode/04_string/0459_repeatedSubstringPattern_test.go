package _4_string

import (
	"fmt"
	"strings"
	"testing"
)

/*
参考这篇题解：https://leetcode.cn/problems/repeated-substring-pattern/solutions/114572/jian-dan-ming-liao-guan-yu-javaliang-xing-dai-ma-s/
如果一个字符串有重复的话，例如：abab (重复子串是ab) 字符串长度为4，重复子串长度为2。
那么，这个字符串移动子串个长度一定会和原来的字符串相等。（前提是子串的长度小于母串的长度，例如：ab->没有重复子串）
如果将母串进行2次拼接，s=abcabc -> ss=abcabcabcabc 那么，如果母串去掉首尾后 ss1=bcabcabcab 如果包含s
那么证明母串由重复的子串构成。
*/
func repeatedSubstringPattern(s string) bool {
	str := s + s
	return strings.Contains(str[1:len(str)-1], s)
}

func TestRepeatedSubstringPattern(t *testing.T) {
	//fmt.Println(repeatedSubstringPattern("abc"))
	fmt.Println(repeatedSubstringPattern("abcabc"))
}
