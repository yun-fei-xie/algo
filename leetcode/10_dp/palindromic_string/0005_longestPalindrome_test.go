package palindromic_string

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

/*
动态规划
以i,j为左右端点的最长回文子串如何思考：

 1. 如果s[i]==s[j]的话，
    a. s[i+1...j-1]也是回文串，那么dp[i][j] = dp[i+1][j-1] + 2
    b. s[i+1...j-1]不是回文串，那么dp[i][j] = dp[i+1][j-1]

 2. 如果s[i]!=s[j]的话，
    dp[i][j] = max{dp[i+1][j], dp[i][j-1]}

如果将递归函数的返回值修改为bool类型，表示内层状态是否是回文串，使用i,j计算回文串的长度，就可以拿到最大值
*/
func longestPalindrome3(s string) string {
	left, right := -1, -1
	maxLen := -1
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// 只有一个字符或者空字符必然是回文串
		if i >= j {
			if j-i+1 > maxLen {
				left = i
				right = j
			}
			return true
		}
		if dfs(i+1, j-1) && s[i] == s[j] {
			if j-i+1 > maxLen {
				left = i
				right = j
			}
			return true
		}
		//两种情况
		dfs(i, j-1)
		dfs(i+1, j)
		return false
	}
	dfs(0, len(s)-1)
	return s[left : right+1]
}

func max(args ...int) int {
	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

/*
动态规划
*/
func longestPalindrome4(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	// dp[i][j] 表示s[i...j]是不是回文串

	// 处理长度为1的
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxLen := 1
	startIndex := -1
	// 处理长度为2的
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			maxLen = 2
			startIndex = i
		}
	}
	// 处理长度大于3的
	// i依赖于i+1 j依赖于j-1 因此，i需要从大到小进行遍历，j需要从小到大进行遍历
	for i := len(s) - 3; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			subLen := j - i + 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if subLen > maxLen {
					maxLen = subLen
					startIndex = i
				}
			}
		}
	}
	return s[startIndex : startIndex+maxLen]
}

func TestLongestPalindrome(t *testing.T) {

	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("bb"))
	//fmt.Println(longestPalindrome2("babad"))
	//fmt.Println(longestPalindrome3("babad"))
	//fmt.Println(longestPalindrome4("babad"))
	//fmt.Println(longestPalindrome3("b"))
	//fmt.Println(longestPalindrome4("b"))
	fmt.Println(longestPalindrome4("ccc"))

}
