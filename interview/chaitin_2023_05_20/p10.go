package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/assign-cookies/
分饼干
*/
func main10() {
	var t int
	var n, k int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&n, &k)

		// 客户
		custom := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&custom[j])
		}
		// 技术支持
		tech := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&tech[j])
		}
		fmt.Println(Math(tech, custom))
	}

}

func Math(tech []int, cust []int) (ans int) {
	sort.Ints(tech)
	sort.Ints(cust)
	//     5 7 9
	// 3 3 3 3 5
	for i, j := len(tech)-1, len(cust)-1; i >= 0 && j >= 0; {
		if tech[i] >= cust[j] {
			ans++
			i--
			j--
		} else {
			j--
		}
	}
	return ans
}
