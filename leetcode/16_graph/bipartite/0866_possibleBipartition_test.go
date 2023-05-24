package bipartite

/*
886. 可能的二分法
https://leetcode.cn/problems/possible-bipartition/

给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和  bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。

示例 1：

输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
输出：true
解释：group1 [1,4], group2 [2,3]
示例 2：

输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
输出：false
示例 3：

输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
输出：false

方法：这个问题也是一个标准的二分图的判定，和785同一个类型
*/
func possibleBipartition(n int, dislikes [][]int) bool {
	// 编号是1...n 整体平移一下0...n-1
	visited := make([]bool, n)
	color := make([]int, n)
	graph := make([][]int, n)
	for i := 0; i < len(dislikes); i++ {
		people1 := dislikes[i][0]
		people2 := dislikes[i][1]
		graph[people1-1] = append(graph[people1-1], people2-1)
		graph[people2-1] = append(graph[people2-1], people1-1)
	}

	var dfs func(i int, c int) bool
	dfs = func(i int, c int) bool {
		visited[i] = true
		color[i] = c

		for _, node := range graph[i] {
			if visited[node] && color[node] != 1-c {
				return false
			}

			if !visited[node] {
				if !dfs(node, 1-c) {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			if !dfs(i, 0) {
				return false
			}
		}
	}
	return true
}
