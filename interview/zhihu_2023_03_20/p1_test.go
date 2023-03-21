package zhihu

import (
	"fmt"
	"testing"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 重建二叉树
 * @param data int整型二维数组 代表节点 id、父节点 id、从左向右的排序信息的二维数组
 * @return int整型一维数组
 */

/*

[[456,123,-1],[789,123,1],[123,-1,0]]
中间为-1 就是根节点
节点id 肯定不会重复 可以用来表示key

mp 递归

*/

func retrieveBinTree(data [][]int) []int {
	// write code here
	tree := map[int]*[2]int{} // 节点id-> [左孩子节点、右孩子节点]
	rootId := -1
	for i := 0; i < len(data); i++ {
		node := data[i]

		nodeId := node[0]
		parentId := node[1]
		position := node[2]

		if parentId == -1 { // 处理根节点
			rootId = nodeId
			//tree[nodeId] = &[2]int{} // 问题就在这里 这里把根的左右子树给还原了 .
			// 这里不需要处理根节点的左右孩子，因为遇到孩子的时候，会处理根节点。 这里只需要记住根节点即可。
			continue
		}

		// 把当前节点放入map
		tree[nodeId] = &[2]int{}

		// 更新自己的父亲节点
		if children, found := tree[parentId]; found {
			if position < 0 {
				(*children)[0] = nodeId
			} else {
				(*children)[1] = nodeId
			}
		} else {
			children := [2]int{}
			if position < 0 {
				children[0] = nodeId
			} else {
				children[1] = nodeId
			}
			tree[parentId] = &children
		}

	}

	var res = make([]int, 0)
	var preOrder func(id int)
	preOrder = func(id int) {
		if children, found := tree[id]; !found {
			return
		} else {
			// 访问左子树
			preOrder(children[0])
			res = append(res, id)
			preOrder(children[1])
			// 访问右子树

		}
	}

	preOrder(rootId)
	return res
}

func TestRetrieveBinTree(t *testing.T) {

	fmt.Println(retrieveBinTree([][]int{{456, 123, -1}, {789, 123, 1}, {123, -1, 0}}))

}

func TestMap(t *testing.T) {
	mp := map[int]*[2]int{}
	mp[0] = &[2]int{1, 2}
	fmt.Println(mp[0])
	key := 3
	if value, ok := mp[key]; ok {
		value[0] = 3
	} else {
		mp[key] = &[2]int{7, 8}
	}
	fmt.Println(mp[key])
}
