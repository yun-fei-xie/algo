package _6_graph

import (
	"fmt"
	"testing"
)

/*
https://leetcode.cn/problems/course-schedule/?favorite=2cktkvj
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。


示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

方法1：通过深度优先遍历判断图中是否有环


方法2：通过拓扑排序判断


*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		in := prerequisites[i][0]
		out := prerequisites[i][1]
		graph[out] = append(graph[out], in)
	}
	var hasCycle bool
	var onPath = make([]bool, numCourses)
	var visited = make([]bool, numCourses)

	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] || hasCycle {
			return
		}
		visited[i] = true
		onPath[i] = true
		for _, node := range graph[i] {
			if !visited[node] {
				dfs(node)
				//可以放在这里判断：如果一个节点已经被访问过，并且在path中，那么肯定形成了环
			} else if visited[node] && onPath[node] {
				hasCycle = true
				return
			}
		}
		onPath[i] = false
	}

	for i := 0; i < numCourses; i++ {
		dfs(i)
	}

	return !hasCycle
}

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
