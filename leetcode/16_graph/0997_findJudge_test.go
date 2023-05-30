package _6_graph_test

/*
997. 找到小镇的法官
https://leetcode.cn/problems/find-the-town-judge/description/?envType=study-plan-v2&envId=graph-theory

小镇里有 n 个人，按从 1 到 n 的顺序编号。传言称，这些人中有一个暗地里是小镇法官。
如果小镇法官真的存在，那么：
小镇法官不会信任任何人。
每个人（除了小镇法官）都信任这位小镇法官。
只有一个人同时满足属性 1 和属性 2 。
给你一个数组 trust ，其中 trust[i] = [ai, bi] 表示编号为 ai 的人信任编号为 bi 的人。
如果小镇法官存在并且可以确定他的身份，请返回该法官的编号；否则，返回 -1 。

方法：图的顶点的度。
1.法官的入度必须是n-1，出度必须是0
2.遍历完trust数组后，法官的入度必须为n-1，出度必须为0，两者的和必须是n-1
*/
func findJudge(n int, trust [][]int) int {
	degree := make([]int, n+1)
	for _, t := range trust {
		degree[t[0]]--
		degree[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if degree[i] == n-1 {
			return i
		}
	}
	return -1
}
