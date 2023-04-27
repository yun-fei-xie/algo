package _4_string

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	arr := make([]int, 0)
	// 321
	n := int(math.Abs(float64(x)))
	for n != 0 {
		arr = append(arr, n%10) // 1->2->3
		n = n / 10
	}
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans = ans*10 + arr[i]
	}
	if x >= 0 && ans <= math.MaxInt32 {
		return ans
	} else if x < 0 && (0-ans) >= math.MinInt32 {
		return 0 - ans
	} else {
		return 0
	}
}

func TestReverse(t *testing.T) {
	fmt.Println(reverse(320))
	fmt.Println(reverse(-320))
	fmt.Println(reverse(123))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2147483648))
}
