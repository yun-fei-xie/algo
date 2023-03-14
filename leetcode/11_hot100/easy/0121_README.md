# 121. 买卖股票的最佳时机


## 解题思路

解题思路参考官方题解。
这个解释很合理，很容易理解。

> 我们来假设自己来购买股票。随着时间的推移，每天我们都可以选择出售股票与否。那么，假设在第 i 天，如果我们要在今天卖股票，那么我们能赚多少钱呢？
显然，如果我们真的在买卖股票，我们肯定会想：如果我是在历史最低点买的股票就好了！太好了，在题目中，我们只要用一个变量记录一个历史最低价格 minprice，我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice。
因此，我们只需要遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？当考虑完所有天数之时，我们就得到了最好的答案。


## 解题代码

```go
func maxProfit(prices []int) int {
	// 只能有一次买卖吗？  只能有一次买卖
	minPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else { // 不是最低价格
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}
```

