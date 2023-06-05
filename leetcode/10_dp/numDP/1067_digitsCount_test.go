package numDP__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
1067. 范围内的数字计数
https://leetcode.cn/problems/digit-count-in-range/description/

给定一个在 0 到 9 之间的整数 d，和两个正整数 low 和 high 分别作为上下界。
返回 d 在 low 和 high 之间的整数中出现的次数，包括边界 low 和 high。
示例 1：

输入：d = 1, low = 1, high = 13
输出：6
解释：
数字 d=1 在 1,10,11,12,13 中出现 6 次。注意 d=1 在数字 11 中出现两次。


方法：数位dp
和233这种类型是一致的
*/

func digitsCount(d int, low int, high int) int {

	sl := strconv.Itoa(low - 1)
	sh := strconv.Itoa(high)

	var f func(s string) int
	f = func(s string) int {

		mem := make([][]int, len(s))
		for k := 0; k < len(s); k++ {
			mem[k] = make([]int, len(s)+1)
			for j := 0; j <= len(s); j++ {
				mem[k][j] = -1
			}
		}

		var numdp func(i int, hasLimit int, numCount int, isNum int) (res int)
		numdp = func(i int, hasLimit int, numCount int, isNum int) (res int) {
			if i == len(s) {
				if isNum == 1 {
					return numCount
				}
				return 0
			}
			// 只缓存hasLimit =false 和 isNum为true的情况
			if hasLimit == 0 && isNum == 1 {
				if mem[i][numCount] != -1 {
					return mem[i][numCount]
				} else {
					defer func() {
						mem[i][numCount] = res
					}()
				}
			}

			// 如果之前没有在选 当前依然不选
			if isNum == 0 {
				res += numdp(i+1, 0, numCount, 0)
			}

			// 之前不管选了还是没有选，当前必须选择一下
			// 如果之前都没有选，当前就不能选前导0
			var lower = 0
			var upper = 9

			if hasLimit == 1 {
				upper = int(s[i] - '0')
			}
			if isNum == 0 {
				lower = 1
			}

			for j := lower; j <= upper; j++ {
				var limit int
				if j == upper && hasLimit == 1 {
					limit = 1
				}
				if j == d {
					res += numdp(i+1, limit, numCount+1, 1)
				} else {
					res += numdp(i+1, limit, numCount, 1)
				}
			}
			return res
		}
		return numdp(0, 1, 0, 0)
	}
	return f(sh) - f(sl)
}

func countD(n int, d int) (res int) {
	for n != 0 {
		if n%10 == d {
			res++
		}
		n = n / 10
	}
	return res
}

func TestDigitsCount(t *testing.T) {
	fmt.Println(digitsCount(1, 1, 13))
	fmt.Println(digitsCount(3, 100, 250))
	fmt.Println(digitsCount(0, 1, 10))
}
