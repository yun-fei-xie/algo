package _4_string

import (
	"fmt"
	"testing"
)

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
haystack = "sadbutsad", needle = "sad"

*/

/*
方法1，暴力匹配
*/
func strStr(haystack string, needle string) int {

	m := len(haystack)
	n := len(needle)

	for i := 0; i <= (m - n); i++ { // 枚举haystack字符串中的每一个可能的起始点i (i<=(m-n) 可以提前退出)
		// 每一轮模式串都从0开始，文本串从i开始
		hStartIndex := i
		nStartIndex := 0

		for nStartIndex < n && haystack[hStartIndex] == needle[nStartIndex] {
			hStartIndex++
			nStartIndex++
		}
		if nStartIndex == n {
			return i
		}
	}

	return -1
}

/*
使用rabin-karp算法实现这个任务
*/
func strStr2(haystack string, needle string) int {

	m := len(haystack)
	n := len(needle)
	//hashEncode [a...z]->[0...25]编码
	var hashEncode func(s string) int
	hashEncode = func(s string) int {
		var h int
		for i := 0; i < len(s); i++ {
			h = h*26 + int(s[i]-'a')
		}
		return h
	}
	patternHash := hashEncode(needle)
	for i := 0; i <= (m - n); i++ {
		//text中的每个子串为：text[i:n+i]
		if hashEncode(haystack[i:n+i]) == patternHash {
			return i
		}
	}
	return -1
}

func TestStrstr(t *testing.T) {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr2("sadbutsad", "sad"))
}
