package _2

/*
*
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。

假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。

例如，传入节点 c（位于单向链表 a->b->c->d->e->f 中），将其删除后，剩余链表为 a->b->d->e->f

示例：

输入：节点 5 （位于单向链表 4->5->1->9 中）
输出：不返回任何数据，从链表中删除传入的节点 5，使链表变为 4->1->9

我不知道该传入节点的前一个节点，怎么办？
可以用val覆盖的方式！把val向前覆盖！
*/
func deleteNode1(node *ListNode) {
	var cur = node

	for cur.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	// 最后一个节点赋值为null
	cur = nil
}

/*
*
我需要让当前这个指针，变成当前指针的next指针
于是需要修改指针的内容
*/
func deleteNode2(node *ListNode) {
	*node = *node.Next
}
