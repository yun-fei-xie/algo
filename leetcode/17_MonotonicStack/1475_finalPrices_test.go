package _7_MonotonicStack

import (
	"container/list"
	"fmt"
	"testing"
)

/*
1475. 商品折扣后的最终价格
https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/description/

给你一个数组 prices ，其中 prices[i] 是商店里第 i 件商品的价格。
商店里正在进行促销活动，如果你要买第 i 件商品，那么你可以得到与 prices[j] 相等的折扣，其中 j 是满足 j > i 且 prices[j] <= prices[i] 的 最小下标 ，如果没有满足条件的 j ，你将没有任何折扣。
请你返回一个数组，数组中第 i 个元素是折扣后你购买商品 i 最终需要支付的价格。
方法1：暴力求解

方法2：单调栈
*/
func finalPrices1(prices []int) []int {
	length := len(prices)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = prices[i]
		for j := i + 1; j < length; j++ {
			if prices[j] <= prices[i] {
				ans[i] -= prices[j]
				break
			}
		}
	}
	return ans
}

/*
单调栈
在使用单调栈时，一般在栈中存储元素的下标
prices->[8,4,6,2,3]
对于prices数组来说，对于每一个prices[i]需要求出该元素的右边第一个小于等于prices[i]的元素。
对于3来说，其右侧没有元素，因此arr[4] = -1 ，然后将3压入栈中 stack->[3->]
对于2来说，栈中元素都大于2，依次出栈，最后栈为空，因此，arr[3]=-1 ,然后将2压入栈中，stack->[2->]
对于6来说，栈中元素小于6，因此arr[2] = 2, 然后将6压入栈中，stack->[2->6->]
对于4来说，栈顶元素6大于4，出栈。arr[1]= 2 ，然后将4压入栈中，stack->[2->4->]
对于8来说，栈顶元素4小于8，arr[0] = 4
*/
func finalPrices(prices []int) []int {

	length := len(prices)
	ans := make([]int, length)
	stack := list.New()

	for i := length - 1; i >= 0; i-- {

		for stack.Len() != 0 {
			if stack.Back().Value.(int) > prices[i] {
				stack.Remove(stack.Back())
			} else {
				break
			}
		}

		if stack.Len() == 0 {
			ans[i] = -1
			stack.PushBack(prices[i])
		} else {
			ans[i] = stack.Back().Value.(int)
			stack.PushBack(prices[i])
		}
	}
	return ans
}

func TestFinalPrices(t *testing.T) {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
