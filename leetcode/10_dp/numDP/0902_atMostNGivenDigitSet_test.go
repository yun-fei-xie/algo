package numDP__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
902. 最大为 N 的数字组合
给定一个按 非递减顺序 排列的数字数组 digits 。
你可以用任意次数 digits[i] 来写的数字。
例如，如果 digits = ['1','3','5']，我们可以写数字，如 '13', '551', 和 '1351315'。
返回 可以生成的小于或等于给定整数 n 的正整数的个数 。

方法：数位dp
这个题和1012几乎是一样的。
*/
func atMostNGivenDigitSet(digits []string, n int) int {
	var s = strconv.Itoa(n)
	var length = len(digits)
	var numdp func(i int, hasLimit bool, isNum bool) (res int)
	numdp = func(i int, hasLimit bool, isNum bool) (res int) {
		if i == len(s) {
			if isNum {
				return 1
			}
			return 0
		}
		// 之前没有选过
		if !isNum {
			// 接着不选
			res += numdp(i+1, false, isNum)
		}

		if hasLimit {
			// 找下i这一位可以放置的上界 digits[k]<=int(s[i]-'0') 下界是digits[0]
			for j := 0; j < length; j++ {
				num, _ := strconv.Atoi(digits[j])
				if num > int(s[i]-'0') {
					break
				}
				res += numdp(i+1, num == int(s[i]-'0'), true)
			}
		} else {
			for j := 0; j < length; j++ {
				res += numdp(i+1, false, true)
			}
		}
		return res
	}
	return numdp(0, true, false)
}

/*
记忆化
*/

func atMostNGivenDigitSet2(digits []string, n int) int {
	var s = strconv.Itoa(n)
	var length = len(digits)

	mem := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		mem[i] = []int{-1, -1}
	}

	var numdp func(i int, hasLimit int, isNum bool) (res int)
	numdp = func(i int, hasLimit int, isNum bool) (res int) {
		if i == len(s) {
			if isNum {
				return 1
			}
			return 0
		}

		if isNum {
			if mem[i][hasLimit] != -1 {
				return mem[i][hasLimit]
			} else {
				defer func() {
					mem[i][hasLimit] = res
				}()
			}
		}

		// 之前没有选过
		if !isNum {
			// 接着不选
			res += numdp(i+1, 0, isNum)
		}

		if hasLimit == 1 {
			// 找下i这一位可以放置的上界 digits[k]<=int(s[i]-'0') 下界是digits[0]
			for j := 0; j < length; j++ {
				num, _ := strconv.Atoi(digits[j])
				if num > int(s[i]-'0') {
					break
				}
				var limit int
				if num == int(s[i]-'0') {
					limit = 1
				}

				res += numdp(i+1, limit, true)
			}
		} else {
			for j := 0; j < length; j++ {
				res += numdp(i+1, 0, true)
			}
		}
		return res
	}
	return numdp(0, 1, false)
}

func TestAtMostNGGive(t *testing.T) {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000))
	fmt.Println(atMostNGivenDigitSet([]string{"7"}, 8))
}
