package _0_dp

import (
	"fmt"
	"testing"
)

/*
2466. 统计构造好字符串的方案数

https://leetcode.cn/problems/count-ways-to-build-good-strings/?envType=study-plan-v2&envId=dynamic-programming
给你整数 zero ，one ，low 和 high ，我们从空字符串开始构造一个字符串，每一步执行下面操作中的一种：

将 '0' 在字符串末尾添加 zero  次。
将 '1' 在字符串末尾添加 one 次。
以上操作可以执行任意次。

如果通过以上过程得到一个 长度 在 low 和 high 之间（包含上下边界）的字符串，那么这个字符串我们称为 好 字符串。

请你返回满足以上要求的 不同 好字符串数目。由于答案可能很大，请将结果对 109 + 7 取余 后返回。

示例 1：

输入：low = 3, high = 3, zero = 1, one = 1
输出：8
解释：
一个可能的好字符串是 "011" 。
可以这样构造得到："" -> "0" -> "01" -> "011" 。
从 "000" 到 "111" 之间所有的二进制字符串都是好字符串。
示例 2：

输入：low = 2, high = 3, zero = 1, one = 2
输出：5
解释：好字符串为 "00" ，"11" ，"000" ，"110" 和 "011" 。

方法：枚举操作，每一步都可以选择zero或者是one 两者都是大于0的。
枚举每一步操作，检查字符串长度。如果在区间中，就将ans++
这个想法很直观，也很容易写出代码，但是比较耗时。能不能优化成动态规划。

仔细看，这道题和第70号问题，爬楼梯是一样的
*/
func countGoodStrings(low int, high int, zero int, one int) int {
	var ans int
	var mod = 1000000000 + 7

	var dfs func(strLen int)
	dfs = func(strLen int) {
		if strLen > high {
			return
		}

		if strLen >= low && strLen <= high {
			ans = (ans + 1) % mod
		}
		// 选择0
		dfs(strLen + zero)
		// 选择1
		dfs(strLen + one)
	}
	dfs(0)
	return ans % mod
}

/*
字符串长度不固定。怎么处理？
枚举每一个合法的长度？先做一下，似乎不是最优方案。
*/
func countGoodStrings2(low int, high int, zero int, one int) int {
	var mod = 1000000000 + 7
	mem := make([]int, high+1) //[0...high]
	for i := 0; i < high+1; i++ {
		mem[i] = -1
	}
	// 长度为length的空间可以构造多少种字符串
	var dfs func(length int) int
	dfs = func(length int) int {
		// 恰好用完
		if length == 0 {
			return 1
		}
		// 无法构造成功
		if length < zero && length < one {
			return 0
		}
		if mem[length] != -1 {
			return mem[length]
		}
		var ret int
		defer func() {
			mem[length] = ret
		}()
		// 尝试构造
		if length < zero {
			ret = dfs(length-one) % mod
			return ret
		} else if length < one {
			ret = dfs(length-zero) % mod
			return ret
		} else {
			ret = (dfs(length-one) + dfs(length-zero)) % mod
			return ret
		}
	}
	var ans int
	// 遍历每一个合法的长度
	for i := low; i <= high; i++ {
		ans = (ans + dfs(i)) % mod
	}
	return ans
}

/*
模仿爬楼梯的写法
*/

func countGoodStrings3(low int, high int, zero int, one int) int {
	const mod int = 1e9 + 7
	var dp = make([]int, high+1) //[0...high] dp[i]表示爬i阶楼梯一共有多少种爬法
	dp[0] = 1
	var ans int
	for i := 1; i <= high; i++ {
		if i >= one {
			dp[i] = dp[i] + dp[i-one]%mod
		}
		if i >= zero {
			dp[i] = dp[i] + dp[i-zero]%mod
		}
		if i >= low {
			ans = (ans + dp[i]) % mod
		}
	}
	return ans
}

func TestCountGoodStrings(t *testing.T) {
	fmt.Println(countGoodStrings(3, 3, 1, 1))
	fmt.Println(countGoodStrings2(3, 3, 1, 1))
	fmt.Println(countGoodStrings(2, 3, 1, 2))
	fmt.Println(countGoodStrings2(2, 3, 1, 2))
	fmt.Println(countGoodStrings3(3, 3, 1, 1))
	fmt.Println(countGoodStrings3(2, 3, 1, 2))
}
