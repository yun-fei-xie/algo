package _2_slidingwindow

import (
	"fmt"
	"reflect"
	"testing"
)

/*
https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

比较abc和bca相等的秘诀在于使用两个hash表。
*/
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}

	mp1 := make(map[uint8]int, 0)
	for i := 0; i < len(p); i++ { // 初始化p
		mp1[p[i]]++
	}

	mp2 := make(map[uint8]int, 0)
	left := 0
	right := 0
	for right < len(s) {

		if right < len(p)-1 { // 初始化部分 构建窗口
			mp2[s[right]]++
			right++
		} else { // 窗口滑动部分
			mp2[s[right]]++
			if reflect.DeepEqual(mp1, mp2) {
				res = append(res, left)
			}

			mp2[s[left]]--
			if mp2[s[left]] == 0 {
				delete(mp2, s[left])
			}
			left++
			right++
		}

	}
	return res
}

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
