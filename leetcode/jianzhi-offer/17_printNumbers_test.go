package jianzhi_offer

import (
	"fmt"
	"math"
	"testing"
)

func printNumbers(n int) []int {
	var res = make([]int, 0)
	maxNumber := int(math.Pow10(n))
	for i := 1; i < maxNumber; i++ {
		res = append(res, i)
	}
	return res
}

func TestPrintNumber(t *testing.T) {
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(2))
	fmt.Println(printNumbers(3))
	fmt.Println(printNumbers(4))
}
