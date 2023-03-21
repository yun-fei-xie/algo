package other

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/split-two-strings-to-make-palindrome/
checkPalindromeFormation

给你两个字符串 a 和 b ，它们长度相同。请你选择一个下标，将两个字符串都在 相同的下标 分割开。由 a 可以得到两个字符串： aprefix 和 asuffix ，满足 a = aprefix + asuffix ，同理，由 b 可以得到两个字符串 bprefix 和 bsuffix ，满足 b = bprefix + bsuffix 。请你判断 aprefix + bsuffix 或者 bprefix + asuffix 能否构成回文串。
当你将一个字符串 s 分割成 sprefix 和 ssuffix 时， ssuffix 或者 sprefix 可以为空。比方说， s = "abc" 那么 "" + "abc" ， "a" + "bc" ， "ab" + "c" 和 "abc" + "" 都是合法分割。
如果 能构成回文字符串 ，那么请返回 true，否则返回 false 。
注意， x + y 表示连接字符串 x 和 y 。

输入：a = "ulacfd", b = "jizalu"
输出：true
解释：在下标为 3 处分割：
aprefix = "ula", asuffix = "cfd"
bprefix = "jiz", bsuffix = "alu"
那么 aprefix + bsuffix = "ula" + "alu" = "ulaalu" 是回文串。


*/

func checkPalindromeFormation(a string, b string) bool {

	var check func(s1 string, s2 string) bool
	// 这种检查方法需要分奇数和偶数
	check = func(s1 string, s2 string) bool {

		mid := (len(s1) - 1) / 2
		if len(s1)%2 == 0 {
			mid += 1
		}

		for i := 0; i < mid; i++ {
			if s1[i] != s2[len(s1)-1-i] {
				// 这个题模拟的难点在这个地方，第一次匹配不上的时候，各自尝试自匹配
				if isPalindrome(&s2, i, len(s1)-1-i) == false && isPalindrome(&s1, i, len(s1)-1-i) == false {
					return false
				} else {
					return true
				}

			}
		}
		return true
	}

	return check(a, b) || check(b, a)

}
func isPalindrome(s *string, left, right int) bool {
	for left < right {
		if (*s)[left] != (*s)[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func TestCheckPalindrome(t *testing.T) {
	fmt.Println(checkPalindromeFormation("abdef", "fecab"))
	fmt.Println(checkPalindromeFormation("abda", "acmc"))
	fmt.Println(checkPalindromeFormation("pvhmupgqeltozftlmfjjde", "yjgpzbezspnnpszebzmhvp"))

}
