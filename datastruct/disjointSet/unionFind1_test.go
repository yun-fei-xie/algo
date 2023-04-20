package disjointSet

/*
第一个版本的并查集，
使用一个id数组表示数据。
数组的下标表示数据本身，数组中的值表示数据的groupId

如果两个数据的id值相同，则表示两个数据是互相连接的。
将两个组合并：将其中一个组的id值全部修改为另外一个组的id值。
*/
type unionFind struct {
	id    []int
	count int
}

// 初始时，每个元素各自为一组，互不连接
func initUnionFind(n int) *unionFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &unionFind{
		id:    id,
		count: n,
	}
}

func (un *unionFind) find(p int) int {
	return un.id[p]
}

func (un *unionFind) isConnected(p, q int) bool {
	return un.id[p] == un.id[q]
}

func (un *unionFind) unionElements(p, q int) {
	pId := un.find(p)
	qId := un.find(q)
	if pId == qId {
		return
	}
	for i := 0; i < un.count; i++ {
		if un.id[i] == pId {
			un.id[i] = qId
		}
	}
}
