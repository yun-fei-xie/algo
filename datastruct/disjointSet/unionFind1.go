package disjointSet

/*
第一个版本的并查集，
使用一个id数组表示数据。
数组的下标表示数据本身，数组中的值表示数据的groupId

如果两个数据的id值相同，则表示两个数据是互相连接的。
将两个组合并：将其中一个组的id值全部修改为另外一个组的id值。
*/
type UnionFind struct {
	id    []int
	count int
}

// 初始时，每个元素各自为一组，互不连接
func InitUnionFind(n int) *UnionFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &UnionFind{
		id:    id,
		count: n,
	}
}

func (un *UnionFind) Find(p int) int {
	return un.id[p]
}

func (un *UnionFind) IsConnected(p, q int) bool {
	return un.id[p] == un.id[q]
}

func (un *UnionFind) UnionElements(p, q int) {
	pId := un.Find(p)
	qId := un.Find(q)
	if pId == qId {
		return
	}
	for i := 0; i < un.count; i++ {
		if un.id[i] == pId {
			un.id[i] = qId
		}
	}
}
