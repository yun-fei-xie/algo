package mid

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


典型的拓扑排序
如何表示这个图的关系呢？用邻接矩阵

从入度为0的节点开始进行广度优先遍历，当整个图遍历完成时，如果收入集合中的节点个数
满足题目中要求的数量，则ok.

*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	var res bool

	//graph := make([][]int, 0)
	//visit := make([]bool, numCourses)
	//for i := 0; i < numCourses; i++ {
	//	graph = append(graph, make([]int, numCourses))
	//}
	//
	//for _, pre := range prerequisites {
	//	graph[pre[0]][pre[1]] = 1 // 1代表有边
	//}

	// 从一个入度为0的点开始寻找

	return res
}
