package _0_dp

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

/*
题解区给出的中心扩散的代码写法显然更加优雅，将一个中心点和两个中心点用一个extend函数进行了结合
*/
func countSubstrings2(s string) int {

	var extend func(i, j int) int
	extend = func(i, j int) int {
		var cnt int
		for i >= 0 && j < len(s) && s[i] == s[j] {
			cnt++
			i--
			j++
		}
		return cnt
	}
	var ans int
	for i := 0; i < len(s); i++ {
		ans += extend(i, i)
		ans += extend(i, i+1)
	}
	return ans
}

/*
动态规划解法：
是不是能找到一种递归关系，也就是判断一个子字符串（字符串的下表范围[i,j]）是否回文，依赖于，子字符串（下表范围[i + 1, j - 1]）） 是否是回文。
如果使用dp[i][j]表示s[i...j]是不是回文串，可以使用如下手段缩小问题规模
1.如果j-i==0，也就是i==j，必然是回文
2.如果j-i==1，也就是j与i相邻，如果s[i]==s[j],s[i...j]也是回文。
3.如果j-i>1，那么如果s[i]==s[j] 并且 s[i+1,j-1]也是回文的话，s[i...j]就是回文。

由于dp[i][j] 依赖 dp[i+1][j-1] ,也及时当前状态依赖于左下角的状态。
因此，dp数组的更新顺序应该从下到上，从左到右。
同时，因为i>=j的，所以，dp数组的只需要更新一半。
*/
func countSubStrings3(s string) int {
	lens := len(s)
	dp := make([][]bool, lens)
	for i := 0; i < lens; i++ {
		dp[i] = make([]bool, lens)
	}
	var ans int
	for i := lens - 1; i >= 0; i-- {
		for j := i; j < lens; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					ans++
					dp[i][j] = true
				} else {
					if dp[i+1][j-1] {
						ans++
						dp[i][j] = true
					}
				}
			}
		}
	}
	return ans
}

func TestCountSubString(t *testing.T) {
	//fmt.Println(countSubstrings("aaa"))
	//fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("fdsklf"))
	fmt.Println(countSubstrings2("fdsklf"))
	fmt.Println(countSubStrings3("fdsklf"))
}
