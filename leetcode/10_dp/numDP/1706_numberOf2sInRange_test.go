package numDP__test

import (
	"strconv"
)

/*
面试题 17.06. 2出现的次数
https://leetcode.cn/problems/number-of-2s-in-range-lcci/description/
编写一个方法，计算从 0 到 n (含 n) 中数字 2 出现的次数。

示例:

输入: 25
输出: 9
解释: (2, 12, 20, 21, 22, 23, 24, 25)(注意 22 应该算作两次)

方法：数位dp
这个问题和233一模一样，一个统计1，一个统计2
*/
func numberOf2sInRange(n int) int {
	s := strconv.Itoa(n)
	// mem数组需要记忆从i这个位置开始，[i...len(s))，能够出现多少个1.
	// 同样沿用之前

	mem := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = make([]int, len(s)+1)
		for j := 0; j < len(s)+1; j++ {
			mem[i][j] = -1
		}

	}

	// 递归函数是在构造数字，如果构造成功，它会返回这个数字中1的个数。
	var numDp func(i int, countTwo int, hasLimit bool) (res int)
	numDp = func(i int, countTwo int, hasLimit bool) (res int) {
		if i == len(s) {
			return countTwo
		}

		if !hasLimit {
			if mem[i][countTwo] != -1 {
				return mem[i][countTwo]
			}
			defer func() {
				mem[i][countTwo] = res
			}()
		}

		if hasLimit {
			for j := 0; j <= int(s[i]-'0'); j++ {
				if j == 2 {
					res += numDp(i+1, countTwo+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countTwo, hasLimit && j == int(s[i]-'0'))
				}
			}
		} else {
			for j := 0; j <= 9; j++ {
				if j == 2 {
					res += numDp(i+1, countTwo+1, hasLimit && j == int(s[i]-'0'))
				} else {
					res += numDp(i+1, countTwo, hasLimit && j == int(s[i]-'0'))
				}
			}
		}
		return res
	}
	return numDp(0, 0, true)
}
