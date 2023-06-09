# 树状数组

Binary Index Tree, BIT

## 解决区间查询问题

```text
| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 |
      1   7   3   0   5   8   3   2   6   2     1   1     4    5
```

BIT巧妙地通过数组的索引构建了一棵多叉树。

## 数组索引和节点之间的关系

BIT使用一个虚拟节点0作为整棵树的根节点。
它的一层孩子节点是索引的二进制位只有1一个1的节点。例如在本例中：{0001,0010,0100,1000} 翻译成10进制就是{1,2,4,8}
这里没有16是因为16越界。
每个子节点要比父亲节点的二进制位多一个1，并且这个1只能通过变换父亲节点高位1左侧的0实现。

1. 什么叫高位1左侧的0？对于上面的第一层节点来说，二进制可以这样写：{0001,0010,0100,1000}高位的前导0可以抹掉。于是
   就成了{1,10,100,1000}。对于1来说，它没有左侧的0，所以它没有孩子。对于4(100)来说，它的左侧有2个0，于是可以进行变换，从低位开始翻转1位0。
   那么节点4就有2个孩子，从左到右的顺序是{101}->5 {110}->6
2. 通过 remove lowbit(1)实现当前节点找父亲节点。
   例如对于索引6{110}来说，移除它的lowbit(1)之后，就比那成了{100}，也就是节点4。移除的方式如下： 假设节点的二进制表示形式是x,那么
   parent(x) = x - (x & (-x))

## 如何更新值

更新只需要更新当前节点和当前节点右边的兄弟节点

通过remove lowbit(1) 可以实现找父亲节点




## 如何查找








## 区间求和问题

一道来自leetcode的题目，区间求和 可修改值
https://leetcode.cn/problems/range-sum-query-mutable/description/

```text
type NumArray struct {
	bit []int
	num []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	bit := make([]int, length+1)
	for i := 0; i < length; i++ {
		update(bit, i+1, nums[i])
	}
	return NumArray{bit: bit, num: nums}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.num[index]
	this.num[index] = val // 每次记住修改nums[index]
	update(this.bit, index+1, diff)
}
func update(bit []int, index int, val int) {
	for index < len(bit) {
		bit[index] += val
		index += index & (-index)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.getSum(right+1) - this.getSum(left)
}

func (this *NumArray) getSum(index int) (ans int) {
	for index != 0 {
		ans += this.bit[index]
		index -= index & (-index)
	}
	return ans
}
```




