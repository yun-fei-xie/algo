package _6_graph_test

/*
1557. 可以到达所有点的最少点数目
https://leetcode.cn/problems/minimum-number-of-vertices-to-reach-all-nodes/description/?envType=study-plan-v2&envId=graph-theory

给你一个 有向无环图 ， n 个节点编号为 0 到 n-1 ，以及一个边数组 edges ，其中 edges[i] = [fromi, toi] 表示一条从点  fromi 到点 toi 的有向边。
找到最小的点集使得从这些点出发能到达图中所有点。题目保证解存在且唯一。
你可以以任意顺序返回这些节点编号。

方法：收集所有度为0的顶点，就是最后的答案。
如何理解？
1. 假设入度为0的顶点集合是s=[v1,v2,v3]
2. 题目给定的图本身是一个连通的图。
3. 如果去掉s中的任意一个顶点，那么因为这个被去掉的顶点入度为0，没有其他边进入，那么在后续的图的遍历中一定不能从其他点到达。
4. 那么，这就是最小集合吗？假设v4的入度不为0，是否一定可以从[v1、v2、v3]到达？ 反证法：假设v4不可达，那么v4的入度必然是0。

*/

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegree := make([]int, n)

	for _, edge := range edges {
		inDegree[edge[1]]++
	}
	var ans = make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans

}
