package mid

import (
	"fmt"
	"testing"
)

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
回文字符串 是正着读和倒过来读一样的字符串。
子字符串 是字符串中的由连续字符组成的一个序列。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"

思路：中心扩散，和第五题最长回文子串相同的解法

*/

func countSubstrings(s string) int {
	res := len(s) // 单个字符

	for i := 0; i < len(s); i++ {
		// 考虑奇数

		for left, right := i-1, i+1; left >= 0 && right < len(s) && s[left] == s[right]; {
			res++
			left--
			right++
		}

		// 考虑偶数 i 与 i-1
		for left, right := i-1, i; left >= 0 && right < len(s) && s[left] == s[right]; {
			res++
			left--
			right++
		}
	}

	return res
}

func TestCountSubString(t *testing.T) {
	//fmt.Println(countSubstrings("aaa"))
	//fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("fdsklf"))
}
