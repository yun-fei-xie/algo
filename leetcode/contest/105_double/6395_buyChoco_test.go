package _105_double

import (
	"fmt"
	"sort"
	"testing"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if len(prices) < 2 || prices[0]+prices[1] > money {
		return money
	}

	return money - (prices[0] + prices[1])
}

func TestBuyChoco(t *testing.T) {
	fmt.Println(buyChoco([]int{1, 2, 2}, 3))
	fmt.Println(buyChoco([]int{3, 2, 3}, 3))
}
