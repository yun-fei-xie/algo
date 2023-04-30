package jianzhi_offer

import (
	"fmt"
	"testing"
)

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。
但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/

/*
可以理解为带有障碍的遍历问题。条件限制就是障碍
*/
func movingCount(m int, n int, k int) int {
	var used = make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var ans int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !outOfBound(m, n, x, y) && used[x][y] == false && valid(x, y, k) {
			ans++
			used[x][y] = true
			dfs(x, y+1)
			dfs(x, y-1)
			dfs(x+1, y)
			dfs(x-1, y)
		}
	}
	dfs(0, 0)
	return ans
}

func valid(x, y, k int) bool {

	var bitSum int
	for x != 0 {
		bitSum += x % 10
		x = x / 10
	}
	for y != 0 {
		bitSum += y % 10
		y = y / 10
	}

	if bitSum <= k {
		return true
	}
	return false
}

func TestMovingCount(t *testing.T) {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}
