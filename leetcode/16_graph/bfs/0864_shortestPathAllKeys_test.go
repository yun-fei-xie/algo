package bfs

import (
	"container/list"
	"fmt"
	"testing"
)

/*
864. 获取所有钥匙的最短路径
https://leetcode.cn/problems/shortest-path-to-get-all-keys/description/

给定一个二维网格 grid ，其中：
'.' 代表一个空房间
'#' 代表一堵墙
'@' 是起点
小写字母代表钥匙
大写字母代表锁
我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们手里有对应的钥匙，否则无法通过锁。
假设 k 为 钥匙/锁 的个数，且满足 1 <= k <= 6，字母表中的前 k 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。
返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回 -1 。

方法：两次遍历。第一次遍历收集钥匙的数量和整个图的起始点。

	第二次使用广度优先遍历，搜素最短路径。

这道题的难点是：
 1. 与常规的简单BFS不同，在这个题中，某个节点访问的节点后续的路径其实还有可能再经过。
    这是因为，每个顶点的状态不只是由它所在位置决定，还有一个很关键的因素就是钥匙。
    所以，visited数组还需要增加一个维度，就是钥匙。
 2. 状态表示。加入整个图中有k把key。那么如何表示k把key都收集完毕呢？可以长度为k的二进制位进行表示。
    例如：如果有3把钥匙，表示3把钥匙都找到可以用一个3位二进制数字表示：111，也就是2^3-1。
    如果有k把钥匙，全部钥匙都取到了可以表示为：2^k-1。  如果给每一把钥匙编号，那么如何判断第i把钥匙
    已经被找到了呢？比如判断从左到右第二位的数字是否为1，可以将一个数字的第二位置为1（其他位为0）然后与钥匙的状态做与运算，然后看结果是否为1。
    010001
    010000 将1左移i-1位，然后与状态进行与运算。如果结果为1，则钥匙在第i位上的值就是1。否则就是0。
    相应的，如果要将钥匙第i位的状态置为1，只需要将与运算改为或运算。

其他写法：
 1. 方向数组 dx=[-1 ,1 , 0 , 0] dy = [0 ,0 , -1 , 1] 方向按照上下左右的顺序。当前的左边{x ,y}就可以转移成 x+dx[i] y+dy[i]
*/

type state struct {
	x   int
	y   int
	key int
}

func shortestPathAllKeys(grid []string) int {
	row, col := len(grid), len(grid[0])
	keyCount := 0
	var start *state
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				keyCount++
			}
			if grid[i][j] == '@' {
				start = &state{x: i, y: j, key: 0}
			}
		}
	}
	// 方向数组
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// 访问标记数组
	visited := make([][][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([][]bool, col)
		for j := 0; j < col; j++ {
			visited[i][j] = make([]bool, 1<<keyCount) // 假如有2把钥匙，所有的状态就是2^2->{00,10,01,11} 所有钥匙都找到2^2-1
		}
	}
	// 开始广度优先遍历
	queue := list.New()
	queue.PushBack(start)
	visited[start.x][start.y][start.key] = true
	step := 0
	for queue.Len() != 0 {
		qSize := queue.Len()
		for i := 0; i < qSize; i++ {
			vertex := queue.Remove(queue.Front()).(*state)
			fmt.Printf("{step->%d nextX->%d nextY->%d char->%c key->%b} \n", step, vertex.x, vertex.y, grid[vertex.x][vertex.y], vertex.key)
			if vertex.key == 1<<keyCount-1 {
				return step
			}
			for j := 0; j < 4; j++ {
				nextX := vertex.x + dx[j]
				nextY := vertex.y + dy[j]
				if nextX < 0 || nextX >= row || nextY < 0 || nextY >= col {
					continue
				}
				key := vertex.key
				c := grid[nextX][nextY]

				if c == '#' {
					continue
				} else if c == '.' {
					if !visited[nextX][nextY][key] {
						queue.PushBack(&state{nextX, nextY, key})
						visited[nextX][nextY][key] = true
					}
				} else if c >= 'A' && c <= 'Z' {
					if key>>(c-'A')&1 == 1 {
						if !visited[nextX][nextY][key] {
							queue.PushBack(&state{nextX, nextY, key})
							visited[nextX][nextY][key] = true
						}
					}
				} else if c >= 'a' && c <= 'z' {
					key = key | (1 << (c - 'a'))
					if !visited[nextX][nextY][key] {
						queue.PushBack(&state{nextX, nextY, key})
						visited[nextX][nextY][key] = true
					}
				}
			}
		}
		step++
	}
	return -1
}

func TestShortestPathAllKeys(t *testing.T) {
	//fmt.Println(shortestPathAllKeys([]string{
	//	"@.a..",
	//	"###.#",
	//	"b.A.B"}))
	//fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
	fmt.Println(shortestPathAllKeys([]string{
		"@...a",
		".###A",
		"b.BCc"}))
}
