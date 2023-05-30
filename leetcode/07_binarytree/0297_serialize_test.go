package _7_binarytree_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*



方法：不管用哪种遍历方式，最后的目的是获得反序列话的结果，也就是能够还原回来。
现在用前序遍历收集节点，然后再反序列化。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// 示例1中的二叉树经过前序遍历后的序列化会得到 1,2,null,null,3,4,null,null,5,null,null,
func (this *Codec) serialize(root *TreeNode) string {

	var serializeString strings.Builder
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			serializeString.WriteString("null,")
			return
		}
		serializeString.WriteString(strconv.Itoa(node.Val) + ",")
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)
	return serializeString.String()
}

// Deserializes your encoded data to tree.
// 1,2,null,null,3,4,null,null,5,null,null,
// 反序列化
func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	list = list[:len(list)-1]

	var buildTree func(list *[]string) *TreeNode
	buildTree = func(list *[]string) *TreeNode {
		// 取出第一号元素
		root := (*list)[0]
		*list = (*list)[1:]
		if root == "null" {
			return nil
		}
		val, _ := strconv.Atoi(root)
		node := &TreeNode{Val: val}
		node.Left = buildTree(list)
		node.Right = buildTree(list)
		return node
	}
	return buildTree(&list)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func TestSeralizeTree(t *testing.T) {
	s := "1,2,null,null,3,4,null,null,5,null,null,"
	s = s[:len(s)-1]
	ss := strings.Split(s, ",")
	fmt.Println(ss)
	fmt.Println(len(ss))
}
