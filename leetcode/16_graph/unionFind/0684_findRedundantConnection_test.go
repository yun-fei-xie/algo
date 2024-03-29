package unionFind_test

/*
684.冗余连接
https://leetcode.cn/problems/redundant-connection/

树可以看成是一个连通且 无环 的 无向 图。
给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。

方法：并查集
如果当前的添加的边会让已经添加的边形成环。那么这条边就是题目说的冗余的边。
*/
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := InitUnionFind2(n + 1)

	for _, edge := range edges {

		v1 := edge[0]
		v2 := edge[1]
		if uf.isConnected(v1, v2) {
			return edge
		} else {
			uf.unionElements(v1, v2)
		}
	}
	return nil
}
