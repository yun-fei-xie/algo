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

	for i := 0; i <= (m - n); i++ { // 枚举haystack字符串中的每一个可能的起始点i
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

func TestStrstr(t *testing.T) {
	fmt.Println(strStr("sadbutsad", "sad"))
}
