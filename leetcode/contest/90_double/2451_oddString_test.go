package _0_double

import (
	"fmt"
	"testing"
)

/*
2451. 差值数组不同的字符串
https://leetcode.cn/problems/odd-string-difference/

给你一个字符串数组 words ，每一个字符串长度都相同，令所有字符串的长度都为 n 。
每个字符串 words[i] 可以被转化为一个长度为 n - 1 的 差值整数数组 difference[i] ，其中对于 0 <= j <= n - 2 有 difference[i][j] = words[i][j+1] - words[i][j] 。注意两个字母的差值定义为它们在字母表中 位置 之差，也就是说 'a' 的位置是 0 ，'b' 的位置是 1 ，'z' 的位置是 25 。
比方说，字符串 "acb" 的差值整数数组是 [2 - 0, 1 - 2] = [2, -1] 。
words 中所有字符串 除了一个字符串以外 ，其他字符串的差值整数数组都相同。你需要找到那个不同的字符串。
请你返回 words中 差值整数数组 不同的字符串。

方法：求出了差值数组，如何快速找到不一样的。（数组比较？）
1. 分类讨论
2. 通过前两个字符串定位那个不一样的。
3. 如果diff0==diff1，那么不同的字符串肯定在words[2...n-1]中，遍历words[2...n-1]找出第一个diffi!=diff0的字符串
4. 如果diff0!=diff1,那么不同的字符串肯定在diff0和diff1之间，用diff2和diff0进行比较。如果diff2==diff0，那么不同的字符串就是diff1，反之是diff0
*/
func oddString(words []string) string {
	var diff0 []uint8 = make([]uint8, 0)
	var diff1 []uint8 = make([]uint8, 0)
	for j := 1; j < len(words[0]); j++ {
		diff0 = append(diff0, words[0][j]-words[0][j-1])
	}
	for j := 1; j < len(words[0]); j++ {
		diff1 = append(diff1, words[1][j]-words[1][j-1])
	}

	if equalSlice(diff0, diff1) {
		for i := 2; i < len(words); i++ {
			diffi := make([]uint8, 0)
			for j := 1; j < len(words[i]); j++ {
				diffi = append(diffi, words[i][j]-words[i][j-1])
			}
			if !equalSlice(diff0, diffi) {
				return words[i]
			}
		}
	} else {
		diff2 := make([]uint8, 0)
		for j := 1; j < len(words[2]); j++ {
			diff2 = append(diff2, words[2][j]-words[2][j-1])
		}
		if !equalSlice(diff0, diff2) {
			return words[0]
		} else {
			return words[1]
		}
	}
	return ""
}

func equalSlice(diff0 []uint8, diff1 []uint8) bool {
	for i := 0; i < len(diff1); i++ {
		if diff0[i] != diff1[i] {
			return false
		}
	}
	return true
}

func TestOddString(t *testing.T) {
	//fmt.Println(oddString([]string{"adc", "wzy", "abc"}))
	fmt.Println(oddString([]string{"mll", "edd", "jii", "tss", "fee", "dcc", "nmm", "abb", "utt", "zyy", "xww", "tss", "wvv", "xww", "utt"}))
}
