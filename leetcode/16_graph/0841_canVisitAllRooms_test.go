package _6_graph_test

/*
841. 钥匙和房间
https://leetcode.cn/problems/keys-and-rooms/?envType=study-plan-v2&envId=graph-theory

有 n 个房间，房间按从 0 到 n - 1 编号。最初，除 0 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。
当你进入一个房间，你可能会在里面找到一套不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。
给你一个数组 rooms 其中 rooms[i] 是你进入 i 号房间可以获得的钥匙集合。如果能进入 所有 房间返回 true，否则返回 false。

方法：深度优先遍历，探索图的连通性。一开始只有0号房间是没有锁的，所以只能从0开始进行遍历。

	这幅图是有向图。
*/
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for _, room := range rooms[i] {
			if !visited[room] {
				dfs(room)
			}
		}
	}
	dfs(0)
	for i := 0; i < n; i++ {
		if visited[i] == false {
			return false
		}
	}
	return true
}
