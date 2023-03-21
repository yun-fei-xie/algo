package main

import (
	"fmt"
)

/*
时间限制： 3000MS
内存限制： 589824KB
题目描述：
小美在玩一项游戏。该游戏的目标是尽可能抓获敌人。

敌人的位置将被一个二维坐标 (x, y) 所描述。

小美有一个全屏技能，该技能能一次性将若干敌人一次性捕获。

捕获的敌人之间的横坐标的最大差值不能大于A，纵坐标的最大差值不能大于B。

现在给出所有敌人的坐标，你的任务是计算小美一次性最多能使用技能捕获多少敌人。

第一行三个整数N,A,B，表示共有N个敌人，小美的全屏技能的参数A和参数B。

接下来N行，每行两个数字x,y，描述一个敌人所在的坐标。

1≤N≤500，1≤A,B≤1000，1≤x,y≤1000。

一行，一个整数表示小美使用技能单次所可以捕获的最多数量。

这种搜索方式估计超时了
*/
func main1() {

	var n, a, b int
	fmt.Scanln(&n, &a, &b)

	var rowMax, colMax = 0, 0

	var e = [1001][1001]int{} // 二维数组
	for i := 0; i < n; i++ {
		var t = [2]int{}
		fmt.Scanln(&t[0], &t[1])
		if t[0] > rowMax {
			rowMax = t[0]
		}
		if t[1] > colMax {
			colMax = t[1]
		}

		e[t[0]][t[1]] = 1 // 这个位置有一个敌人
	}

	var max = 0

	for i := 1; i <= rowMax; i++ {

		for j := 1; j <= colMax; j++ { // 站在每一个点 计算辐射范围内的敌人
			cnt := calc(&e, i, i+a, j, j+b)
			if cnt > max {
				max = cnt
			}
		}

	}

	fmt.Println(max)

}

func calc(matrix *[1001][1001]int, left, right, up, down int) int {
	sum := 0
	for i := left; i < len(matrix) && i <= right; i++ {
		for j := up; j < len(matrix) && j <= down; j++ {
			if matrix[i][j] == 1 {
				sum++
			}
		}
	}
	return sum
}
