# leetcode 59. 螺旋矩阵 II

## 题目链接

https://leetcode.cn/problems/spiral-matrix-ii/

## 题目描述

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

## 解题思路
1. 生成符合大小的二维数组
2. 从1~n^2把数字一步一步填充到数组中。
难点在第二步，如何填充。按照约定的方向，顺时针填充。

这里面最重要的是4个变量
```go
// 行的left、right 其实是列号
	// 列的left、right 其实是行号
	rowLeft := 0
	rowRight := n - 1
	colLeft := 0
	colRight := n - 1
```
他们控制着填充的范围。比如，在填充第一行的时候，行号是固定的，即rowLeft = 0。
此时变化的是列下标i,那么这个i的范围在哪里呢？答案是：[colLeft , colRight]。
而这一行填充完毕后，行的范围就减少了-> rowLeft ++ 
其他三个方向依此类推。

如何按照顺序进行逆时针旋转-> 使用一个step数组进行标记。
``step[stepIndex%len(step)] == 1``
1->2->3->4->1->2->3->4 ... 这样一直旋转。


## 解题代码

```go
func generateMatrix(n int) [][]int {
	// init
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 行的left、right 其实是列号
	// 列的left、right 其实是行号
	rowLeft := 0
	rowRight := n - 1
	colLeft := 0
	colRight := n - 1

	step := [4]int{1, 2, 3, 4} // 四步走战略
	stepIndex := 0

	for num := 1; num <= n*n; {
		if step[stepIndex%len(step)] == 1 {
			for i := colLeft; i <= colRight; i++ {
				res[rowLeft][i] = num // 行下标不变 移动列下标
				num++
			}
			rowLeft++
			stepIndex++
		} else if step[stepIndex%len(step)] == 2 {
			for i := rowLeft; i <= rowRight; i++ {
				res[i][colRight] = num // 列下标不变 移动行下标
				num++
			}
			colRight--
			stepIndex++
		} else if step[stepIndex%len(step)] == 3 {
			for i := colRight; i >= colLeft; i-- {
				res[rowRight][i] = num
				num++
			}
			rowRight--
			stepIndex++
		} else if step[stepIndex%len(step)] == 4 {
			for i := rowRight; i >= rowLeft; i-- {
				res[i][colLeft] = num
				num++
			}
			colLeft++
			stepIndex++
		}
	}
	return res
}

func TestGenerateMatrix(t *testing.T) {
	n := 3
	res := generateMatrix(n)

	for i := 0; i < n; i++ {
		fmt.Print("[ ")
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Printf("]\n")
	}
}



```

   