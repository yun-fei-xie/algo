package main

import (
	"fmt"
	"math"
	"sort"
)

/*
5
1 2 3 4 5

有多少子集，满足子集内元素两两之间互为倍数 每次放入集合中的元素需要和前面的元素都构成倍数（先排序，这样后面的添加的元素只需要和上一个元素进行比较）


只能通过20%

*/

func main3() {
	var cnt int
	fmt.Scanln(&cnt)
	arr := make([]int, 0)
	for i := 0; i < cnt; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	sort.Ints(arr)

	path := make([]int, 0)
	var res int64 = 0
	var dfs func(arr *[]int, startIndex int)
	dfs = func(arr *[]int, startIndex int) {
		if startIndex >= len(*arr) { // 递归到底
			if len(path) >= 2 {
				res++
			}
			return
		}

		// 当前元素可放可不放
		// 不放进去
		dfs(arr, startIndex+1)

		// 尝试放进去
		if len(path) == 0 {
			path = append(path, (*arr)[startIndex])
			dfs(arr, startIndex+1)
			path = path[0 : len(path)-1]
		} else {
			pre := path[len(path)-1]
			if (*arr)[startIndex]%pre == 0 {
				path = append(path, (*arr)[startIndex])
				dfs(arr, startIndex+1)
				path = path[0 : len(path)-1]
			}
		}

	}
	dfs(&arr, 0)

	fmt.Printf("%d", int64(res)%int64(math.Pow10(9)+7))

}
