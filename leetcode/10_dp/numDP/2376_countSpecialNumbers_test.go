package numDP__test

import "strconv"

/*
2376. 统计特殊整数
https://leetcode.cn/problems/count-special-integers/description/

如果一个正整数每一个数位都是 互不相同 的，我们称它是 特殊整数 。
给你一个 正 整数 n ，请你返回区间 [1, n] 之间特殊整数的数目。

这道题和1012是一模一样的代码。
方法：数位dp
*/
func countSpecialNumbers(n int) int {

	s := strconv.Itoa(n)
	var f func(s string) int
	f = func(s string) int {
		// 省略一些非常有限的状态，只保存i,mask这两种状态 对于isNum==false 和 hasLimit =true这样的状态直接求解
		// 所以记忆化数组只存保存，isNum==true && hasLimit ==false的状态

		mem := make([][]int, len(s))
		for k := 0; k < len(s); k++ {
			mem[k] = make([]int, 1<<10+1)
			for r := 0; r < 1<<10; r++ {
				mem[k][r] = -1
			}
		}

		var numDP func(i int, mask int, hasLimit bool, isNum bool) (res int)
		numDP = func(i int, mask int, hasLimit bool, isNum bool) (res int) {
			if i == len(s) {
				if isNum {
					return 1
				}
				return res
			}

			if isNum && !hasLimit {
				if mem[i][mask] != -1 {
					return mem[i][mask]
				}
				defer func() {
					mem[i][mask] = res
				}()
			}

			// 枚举位置i上的数字
			// 如果前面的位置一直都没有选择过，当前位置也可以不选(不选就是前导0) 这种情况单独讨论
			if !isNum {
				res += numDP(i+1, mask, false, false)
			}
			//前面选择过 1.当前位置的下界就是0，没选过当前位置的下界就是1（0的情况单独讨论）
			//当前位置有限制，上界就是int(s[i]-'0')，否则就是9
			var lower = 0
			var upper = 9
			if hasLimit {
				upper = int(s[i] - '0')
			}
			if !isNum {
				lower = 1
			}
			// 当前位可以选择的数字
			for j := lower; j <= upper; j++ {
				if (mask>>j)&1 == 1 {
					continue
				} else {

					res += numDP(i+1, mask|(1<<j), hasLimit && j == upper, true)
				}
			}

			return res
		}
		return numDP(0, 0, true, false)
	}
	return f(s)
}
