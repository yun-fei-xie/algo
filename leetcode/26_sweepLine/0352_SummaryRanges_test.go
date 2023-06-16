package _6_sweepLine

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

/*

352. 将数据流变为多个不相交区间
https://leetcode.cn/problems/data-stream-as-disjoint-intervals/

给你一个由非负整数 a1, a2, ..., an 组成的数据流输入，请你将到目前为止看到的数字总结为不相交的区间列表。
实现 SummaryRanges 类：
SummaryRanges() 使用一个空数据流初始化对象。
void addNum(int val) 向数据流中加入整数 val 。
int[][] getIntervals() 以不相交区间 [starti, endi] 的列表形式返回对数据流中整数的总结。

方法：扫描线 分类讨论。
这里直接参考官方题解：
情况一：如果存在一个区间 [l,r]，它完全包含 val，那么val对已经存在的区间不会产生任何影响。
情况二：如果存在一个区间 [l,r]，它的右边界 r「紧贴着」val，即 r+1=val，那么在加入val之后，该区间会从 [l,r]变为[l,r+1]；
情况三：如果存在一个区间 [l,r]，它的左边界 l「紧贴着」val，即 l−1=val，那么在加入val之后，该区间会从 [l,r]变为[l−1,r]；
情况四：如果情况二和情况三同时成立，那么，情况二和情况三种的两个区间会合并为一个大的区间。
情况五：在上述四种情况均不满足的情况下，val 会单独形成一个新的区间 [val,val]

在扫描线中，需要维护区间的顺序。
并且在当前的算法思路中，我们需要能够快速找到，对于给定的val,找到最大的r，并且r<val。找到最小的l,并且l>val。
这个数据结构就是有序映射，或者叫有序map。
*/

type SummaryRanges struct {
	*redblacktree.Tree
}

func Constructor() SummaryRanges {
	return SummaryRanges{redblacktree.NewWithIntComparator()}
}

func (this *SummaryRanges) AddNum(value int) {
	//考虑情况1，是否存在这样一个区间[l,r] l<=val && val<=r
	interval0, found0 := this.Floor(value)
	if found0 && interval0.Value.(int) >= value {
		return
	}

	// 找到interval1=[l1,r1] 并且l1是满足l1>val中的最小值
	interval1 := this.Iterator()
	if found0 {
		interval1 = this.IteratorAt(interval0)
	}
	found1 := interval1.Next()

	// 看看能不能和左边的区间融合
	leftSide := found0 && interval0.Value.(int)+1 == value
	rightSide := found1 && interval1.Key().(int)-1 == value

	// 如果左右都可以融合，那么是情况4
	if leftSide && rightSide {
		interval0.Value = interval1.Value().(int)
		this.Remove(interval1.Key())
	} else if leftSide {
		// 如果只有左边->情况3
		interval0.Value = value
	} else if rightSide {
		// 如果只有右边->情况2
		right := interval1.Value().(int)
		left := value
		this.Remove(interval1.Key())
		this.Put(left, right)
	} else {
		this.Put(value, value)
	}

}

func (this *SummaryRanges) GetIntervals() [][]int {
	var ans = make([][]int, 0)
	for it := this.Iterator(); it.Next(); {
		ans = append(ans, []int{it.Key().(int), it.Value().(int)})
	}
	return ans
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
