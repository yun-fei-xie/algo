package _2_linkedList

/*
https://leetcode.cn/problems/copy-list-with-random-pointer/description/
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
返回复制链表的头节点。
用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
你的代码 只 接受原链表的头节点 head 作为传入参数。

输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
方法1：递归创建，使用一个hash表记录原节点和copy节点之间的创建关系
*/
func copyRandomList(head *Node) *Node {
	mem := make(map[*Node]*Node) // 原节点->copy节点
	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, found := mem[node]; found {
			return n
		}

		// copy 当前节点
		newNode := &Node{Val: node.Val}
		mem[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		// mem[node] = newNode 放在这里会出现 节点以为没有创建过 于是递归创建，形成环。
		return newNode
	}

	return deepCopy(head)
}

/*
方法2：复制

*/

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 节点1:1复制  a->a'->b->b'->c->c'->null
	cur := head
	for cur != nil {
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}

	// 处理x'的随机指针
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	// 将x'节点从复制节点中挑出来
	//  a->a'->b->b'->c->c'->null
	// 让a->b->c->null  a'->b'->c'->null
	// 修改指针 交替修改指针很绕
	//cpHead := head.Next
	//for node := head; node != nil; {
	//	cpNode := node.Next
	//	node.Next = node.Next.Next
	//
	//	cpNode.Next = cpNode.Next.Next
	//}

	dummy := &Node{Val: -1, Next: nil}
	tail := dummy
	// 把x'节点单独拿出来处理，使用尾插法构建新链表比较容易
	for node := head; node != nil; node = node.Next {
		cpNode := node.Next
		node.Next = node.Next.Next
		tail.Next = cpNode
		tail = tail.Next
		tail.Next = nil
	}

	return dummy.Next
}
