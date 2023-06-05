package numDP__test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
1012. 至少有 1 位重复的数字

方法：数位dp + 集合论|二进制枚举
1.最大值是n，将n转换为字符串之后可以拿到字符串的长度。然后从左到右枚举字符串的每一位。
2.我怎么知道，当前选中的这一位数字和之前选中的数字是否构成重复？

方法：正难则反
1. 题目要求至少重复一位，那么可以求无重复的数字个数。然后用总的组合-无重复的组合。

010 是否合法？ 其实合法 010就是10，所以需要记录之前是否填写过数字
*/
func numDupDigitsAtMostN(n int) int {
	// [1...n]
	s := strconv.Itoa(n)
	var f func(s string) int
	f = func(s string) int {
		var numDP func(i int, mask int, hasLimit bool, isNum bool) (res int)
		numDP = func(i int, mask int, hasLimit bool, isNum bool) (res int) {
			if i == len(s) {
				if isNum {
					return 1
				}
				return res
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

	// 总的组合 n (1到n总的组合就是n ,因为1-n一共就n个数字)
	return n - f(s)
}

/*
搭配记忆化
*/
func numDupDigitsAtMostN2(n int) int {

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

	// 总的组合 n (1到n总的组合就是n ,因为1-n一共就n个数字)
	return n - f(s)
}

func TestNumDupDigitsAtMostN(t *testing.T) {
	fmt.Println(numDupDigitsAtMostN(100))
	fmt.Println(numDupDigitsAtMostN(1000))
	fmt.Println(numDupDigitsAtMostN2(100))
	fmt.Println(numDupDigitsAtMostN2(1000))
	//mask := 0b100000
	//m := 1 << 5
	//fmt.Printf("%b\n", mask)
	//fmt.Printf("%b\n", m)
	//fmt.Println((mask)&m == 1)
}
func TestBitOperate(t *testing.T) {
	mask := 0b010100
	fmt.Println(mask&(1<<2) == 1)

	// 上面的写法不能检测 mask从右到左索引为2（索引从0开始）的位置是否为1，010100 & 100 -> 会得到4
	// 而是应该(mask>>2 & 1 )
}
