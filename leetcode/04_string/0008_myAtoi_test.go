package _4_string

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

/*
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
对于溢出的处理方式通常可以转换为 INT_MAX 的逆操作。比如判断某数乘以 101010 是否会溢出，那么就把该数和 INT_MAX 除以 101010 进行比较。

"-4193 with words"
*/
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	var ans int
	var sign int = 1

	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '+' {
			sign = 1
		} else if i == 0 && s[i] == '-' {
			sign = -1
		} else if isNum(s[i]) { // 处理数字
			num := s[i] - '0'
			if ans > math.MaxInt32/10 || ans == math.MaxInt32/10 && num > math.MaxInt32%10 {
				return math.MaxInt32
			}
			if ans < math.MinInt32/10 || ans == math.MinInt32/10 && -1*int(num) < math.MinInt32%10 {
				return math.MinInt32
			}
			ans = ans*10 + sign*int(num)

		} else {
			break
		}
	}
	return ans
}

func isNum(c uint8) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func TestMyAtoi(t *testing.T) {

	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("2147483648"))

	//fmt.Println(uint8(9) > math.MaxInt32%10)

}
