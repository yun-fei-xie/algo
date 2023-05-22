package _0_dp

import (
	"fmt"
	"strconv"
	"testing"
)

/*
91. 解码方法
https://leetcode.cn/problems/decode-ways/description/?envType=study-plan-v2&envId=dynamic-programming

一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：

"AAJF" ，将消息分组为 (1 1 10 6)
"KJF" ，将消息分组为 (11 10 6)
注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。

给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。

题目数据保证答案肯定是一个 32 位 的整数。

方法：递归，每次向下递归必须保证当前切分的数字的合法性。
为什么 i>=length 的时候要返回1？
因为找到了一种合法的方式字符串切分完毕。
所以要返回1。如果找不到合法的切分，就不会走到这里。
*/
func numDecodings(s string) int {
	//每次切分只能切分1位或者两位。如果是1位，则必须在[0...9]如果是2位，则必须在[10...26]
	var length = len(s)
	var dfs func(i int) int // [i...length-1]有多少种解码方式
	dfs = func(i int) int {
		if i >= length {
			return 1
		}
		var ret int
		if s[i] >= '1' && s[i] <= '9' {
			ret += dfs(i + 1)
		}
		if i+1 < length {
			num, _ := strconv.Atoi(s[i : i+2]) //[i...i+1]
			if num >= 10 && num <= 26 {
				ret += dfs(i + 2)
			}
		}
		// 如果找不到合法的切分，这里会返回0
		return ret
	}
	return dfs(0)
}

/*
1:1翻译成递推
*/
func numDecodings2(s string) int {
	length := len(s)
	dp := make([]int, length+1)
	dp[length] = 1
	for i := length - 1; i >= 0; i-- {
		var cnt int
		// 按照题目的输入，单个字符如果不是0就符合要求
		if s[i] != '0' {
			cnt += dp[i+1]
		}
		if i+1 < length {
			num, _ := strconv.Atoi(s[i : i+2])
			if num >= 10 && num <= 26 {
				cnt += dp[i+2]
			}
		}
		dp[i] = cnt
	}
	return dp[0]
}

func TestNumDecodings(t *testing.T) {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("06"))
	fmt.Println(numDecodings2("12"))
	fmt.Println(numDecodings2("226"))
	fmt.Println(numDecodings2("06"))
}
