package mid

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/longest-palindromic-substring/?favorite=2cktkvj

给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

这题没那么简单！

1. 暴力枚举：二重循环,不断验证每一个区间中的子串是不是回文串
2. 中心扩散法  https://leetcode.cn/problems/longest-palindromic-substring/solutions/255195/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
每到一个字符，就已该字符为中心向两边进行扩散
3. 动态规划法  https://leetcode.cn/problems/longest-palindromic-substring/solutions/255195/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
*/

/*
中心扩散法需要考虑中心是奇数还是偶数
*/
func longestPalindrome2(s string) string {
	maxLeft := -1
	maxRight := len(s)
	maxLength := 0

	for i := 0; i < len(s); i++ {
		// 考虑中心为奇数
		for left, right := i, i; left >= 0 && right < len(s) && s[left] == s[right]; { // 注意边界的判断应该放在取字符的前面，否则会出现数组下标越界
			if right-left+1 > maxLength {
				maxLength = right - left + 1
				maxLeft = left
				maxRight = right
			}
			left--
			right++
		}

		//考虑中心为偶数 每次都尝试和左边进行搭配（为什么不尝试和右边进行搭配？因为i+1时，和左边进行搭配就相当于当前和右边进行搭配。）
		for left, right := i-1, i; left >= 0 && right < len(s) && s[left] == s[right]; {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
				maxLeft = left
				maxRight = right
			}
			left--
			right++
		}
	}

	return s[maxLeft : maxRight+1]
}

/*
暴力枚举
*/
func longestPalindrome(s string) string {

	maxLength := 1
	left := 0
	right := 0
	for i := 0; i < len(s); i++ { // 以i为起点
		for j := i; j < len(s); j++ { // 以j为终点 考察 [i...j] 这个区间的子串是不是回文串
			if j-i+1 > maxLength && isPali(&s, i, j) {
				maxLength = j - i + 1
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

func isPali(s *string, left int, right int) bool {

	for left <= right {
		if (*s)[left] != (*s)[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {

	//fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome2("babad"))

}
