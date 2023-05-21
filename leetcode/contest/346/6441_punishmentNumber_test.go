package _46

import (
	"fmt"
	"strconv"
	"testing"
)

func punishmentNumber(n int) int {
	var ans int
	for i := 1; i <= n; i++ {
		if valid(i) {
			ans += i * i
		}
	}
	return ans
}

func valid(n int) bool {
	s := strconv.Itoa(n * n)
	var ans bool
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if i == len(s) {
			if sum == n {
				ans = true
			}
			return
		}
		// 从i这一位开始，尝试取s[i,i] s[i,i+1]...s[i]
		for j := i; j < len(s); j++ {
			num, _ := strconv.Atoi(s[i : j+1])
			dfs(j+1, sum+num)
		}
	}
	dfs(0, 0)
	return ans
}

func TestPunishmentNumber(t *testing.T) {
	fmt.Println()
	fmt.Println(punishmentNumber(10))
	fmt.Println(punishmentNumber(37))
}
