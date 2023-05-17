package sectionDP

/*
受到了375题getMoney的影响。
如何处理合并后的数组长度变小
*/
func mergeStones(stones []int, k int) int {
	length := len(stones)
	if length%(k-1) != 0 {
		return -1
	}
	// 区间快速求和->使用前缀和
	prefix := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			prefix[i] = stones[i]
		}
		prefix[i] = prefix[i-1] + stones[i]
	}

	// 否则就是可以合并的

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j-i+1 < k {
			return 0
		}

		// 合并区间[i...j]，找出最小值 区间长度为k
		// 我是真的想直接融合。但是拼凑出合并的数组开销较大
		return -1
	}
	return dfs(0, len(stones)-1)
}
