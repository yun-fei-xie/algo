package bfs

/*
934. 最短的桥
https://leetcode.cn/problems/shortest-bridge/description/

给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
返回必须翻转的 0 的最小数目。

方法：两次广度优先遍历。第一次遍历找到第一个岛屿中的所有的点。
然后枚举第一座岛屿中的所有点，对每一个点进行广度优先遍历。
第二次bfs的时候，要把第一座岛屿的点排除掉。
*/
func shortestBridge(grid [][]int) (step int) {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			island := []pair{}
			grid[i][j] = -1
			q := []pair{{i, j}}
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				island = append(island, p)
				for _, d := range dirs {
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						q = append(q, pair{x, y})
					}
				}
			}

			q = island
			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return
}
