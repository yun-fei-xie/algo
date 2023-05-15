# leetcode 0051 N皇后

## 题目链接

https://leetcode.cn/problems/n-queens/description/

## 题目描述

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

## 解题思路

这是一个典型的递归回溯问题。
其中的难点在于如何快速判断当前要放的位置(x , y  )是否合法。
如果将二维矩阵看做一个原点在左上角的坐标系，行为x轴，列为y轴
那么使用一维数组 row[] , 一维数组 col[]可以快速判断当前的位置在 行和列上是否合法。

对于撇（丿）这样的对角线，可以看出，处于同一条对角线上的坐标（x + y = k ）是一个固定值,并且这样的对角线有2*n-1 条。(下标从0->
2n-2)（需要有2*n-1位置的数组 声明数组长度的时候要小心）
对于捺（）这样的对角线，可以看出，对于同一条对角线上的坐标（x - y + n -1）是一个固定值，并且这样的对角线也有2*n-2条。(下标从0->
2n-2)
于是就可以用4个bool数组进行合法性检查。

对于满足条件的坐标，可以在遍历时加入到path路径上去，并标记因为该位置的加入影响到的行、列、对角线。
同时进入下一轮递归。 在递归回退时，将当前轮限制的条件去掉。

递归终止条件：
index ==n -> 相当于走到了一个空的位置（此时认为n个节点组成的路径已经搜集齐了）
将当前路径保存起来，并进行回退。

## 解题启发

对于二维平面的搜索问题，在某些情况下建立坐标系可能会对解题有帮助。该坐标系的原点一般是矩阵的左上角（0 ， 0 ）。




## 解题code

```go
func solveNQueens(n int) [][]string {

	res := make([][]string, 0)
	position := make([]int, 0) // 存放路径

	row := make([]bool, n)
	col := make([]bool, n)
	dail1 := make([]bool, 2*n-1) // pie 2*n-1 而不是2*(n-1)
	dail2 := make([]bool, 2*n-1) // nai

	var putQueue func(index int) // index 表示行号
	putQueue = func(index int) {
		if index == n {
			res = append(res, generate(position, n))
			return
		}

		for i := 0; i < n; i++ { //i表示列号
			if !row[i] && !col[i] && !dail1[index+i] && !dail2[index-i+n-1] {
				row[i] = true
				col[i] = true
				dail1[index+i] = true
				dail2[index-i+n-1] = true

				position = append(position, i)
				putQueue(index + 1)
				position = position[:len(position)-1]

				row[i] = false
				col[i] = false
				dail1[index+i] = false
				dail2[index-i+n-1] = false
			}
		}
	}

	putQueue(0)

	return res
}
func generate(position []int, n int) []string {
	res := make([]string, 0)
	for i := 0; i < len(position); i++ {
		sb := strings.Builder{}
		for j := 0; j < n; j++ {
			if j != position[i] {
				sb.WriteString(".")
			} else {
				sb.WriteString("Q")
			}
		}
		res = append(res, sb.String())
	}
	return res
}
```