package disjointSet

/*
重新表示一个组：
将每一个元素都看做一个节点，每个节点指向其父亲节点
使用一个parent数组
0，1，2，3，4，5，6，7
0，1，2，3，4，5，6，7

相对于unionFind1
这种表示方法使得isConnected判断变慢了。
因为每个节点只记录自己的父亲节点。
查找的时间复杂度取决的树的高度。
*/
type unionFind2 struct {
	parent []int
	count  int
}

func initUnionFind2(n int) *unionFind2 {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 初始时，每一个节点的父节点都是本身
	}
	return &unionFind2{
		parent: parent,
		count:  n,
	}
}

func (un *unionFind2) unionElements(p, q int) {
	pRoot := un.find(p)
	qRoot := un.find(q)
	if pRoot == qRoot {
		return
	} else {
		un.parent[pRoot] = qRoot
	}
}

func (un *unionFind2) isConnected(p, q int) bool {
	pRoot := un.find(p)
	qRoot := un.find(q)
	return pRoot == qRoot
}

/*
查找节点p的root节点
*/
//func (un *unionFind2) find(p int) int {
//	for un.parent[p] != p {
//		p = un.parent[p]
//	}
//	return p
//}

/*
在查找的过程中实现路径压缩
*/

func (un *unionFind2) find(p int) int {
	parentId := un.parent[p]
	if parentId == p {
		return parentId
	}
	un.parent[p] = un.find(un.parent[p])
	return un.parent[p]
}
