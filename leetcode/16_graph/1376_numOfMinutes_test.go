package _6_graph_test

import (
	"fmt"
	"math"
	"testing"
)

/*
1376. 通知所有员工所需的时间
https://leetcode.cn/problems/time-needed-to-inform-all-employees/?envType=study-plan-v2&envId=graph-theory

公司里有 n 名员工，每个员工的 ID 都是独一无二的，编号从 0 到 n - 1。公司的总负责人通过 headID 进行标识。
在 manager 数组中，每个员工都有一个直属负责人，其中 manager[i] 是第 i 名员工的直属负责人。对于总负责人，manager[headID] = -1。题目保证从属关系可以用树结构显示。
公司总负责人想要向公司所有员工通告一条紧急消息。他将会首先通知他的直属下属们，然后由这些下属通知他们的下属，直到所有的员工都得知这条紧急消息。
第 i 名员工需要 informTime[i] 分钟来通知它的所有直属下属（也就是说在 informTime[i] 分钟后，他的所有直属下属都可以开始传播这一消息）。
返回通知所有员工这一紧急消息所需要的 分钟数 。

输入：n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]
输出：1
解释：id = 2 的员工是公司的总负责人，也是其他所有员工的直属负责人，他需要 1 分钟来通知所有员工。
上图显示了公司员工的树结构。

方法：感觉可以用多叉树做
*/
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	graph := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		graph[i] = map[int]int{}
	}

	for i := 0; i < n; i++ {
		in := i
		out := manager[i]
		if out == -1 {
			continue
		}
		t := informTime[out]
		graph[out][in] = t
	}

	visited := make([]bool, n)
	var ans = math.MinInt
	var dfs func(i int, cost int)
	dfs = func(i int, cost int) {
		if len(graph[i]) == 0 {
			if cost > ans {
				ans = cost
			}
			return
		}
		for node, time := range graph[i] {
			if !visited[node] {
				dfs(node, cost+time)
			}
		}
	}
	dfs(headID, 0)
	return ans

}

func TestNumOfMinutes(t *testing.T) {
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
}
