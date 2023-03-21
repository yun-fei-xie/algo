package main

import (
	"fmt"
	"math"
)

/*
商店
时间限制： 3000MS
内存限制： 589824KB
题目描述：
现在商店里有N个物品，每个物品有原价和折扣价。

小美想要购买商品。小美拥有X元，一共Y张折扣券。

小美需要最大化购买商品的数量，并在所购商品数量尽量多的前提下，尽量减少花费。

你的任务是帮助小美求出最优情况下的商品购买数量和花费的钱数。

输入描述
第一行三个整数，以空格分开，分别表示N,X,Y。

接下来N行，每行两个整数，以空格分开，表示一个的原价和折扣价。

1≤N≤100, 1≤X≤5000, 1≤Y≤50，每个商品原价和折扣价均介于[1,50]之间。

输出描述
一行，两个整数，以空格分开。第一个数字表示最多买几个商品，第二个数字表示在满足商品尽量多的前提下所花费的最少的钱数。

样例输入
3 5 1
4 3
3 1
6 5
样例输出
2 5

提示
样例解释1

第一个商品原价购入，第二个商品折扣价购入，可以获得最多的商品数量2个。

此时消耗5元。因此输出 2 5。

输入样例2

3 5 1

4 3

3 1

6 1

输出样例2

2 4

样例解释2

可以发现有很多种买两个商品的方法。

最省钱的方案是第二个商品原价购入，第三个商品折扣价购入。此时花费4元。

输入样例3

10 30 3

2 1

3 2

2 1

10 8

6 5

4 3

2 1

10 9

5 4

4 2

输出样例3

8 24

ac  36%
*/
func main4() {

	var n, x, y int // n 个商品，x元，y个折扣
	fmt.Scanln(&n, &x, &y)
	goods := make([][2]int, 0)

	minPrice := math.MaxInt
	minDiscount := math.MaxInt

	for i := 0; i < n; i++ {
		temp := [2]int{}
		fmt.Scanln(&temp[0], &temp[1])
		if temp[0] < minPrice {
			minPrice = temp[0]
		}
		if temp[1] < minDiscount {
			minDiscount = temp[1]
		}
		goods = append(goods, temp)
	}

	// 尝试暴力解法 每个商品 要么原价 要么折扣 还可以不买 （三种情况）
	maxCnt := math.MinInt32
	minSpend := math.MaxInt

	var dfs func(arr [][2]int, spend int, discount int, startIndex int, cnt int)
	dfs = func(arr [][2]int, spend int, discount int, startIndex int, cnt int) {
		if startIndex >= len(arr) { // 递归到底
			if cnt >= maxCnt {
				maxCnt = cnt
				if spend < minSpend {
					minSpend = spend
				}
			}
			return
		}

		if x-spend < minDiscount {
			if cnt >= maxCnt {
				maxCnt = cnt
				if spend < minSpend {
					minSpend = spend
				}
			}
			return
		}

		// 尝试原价买
		if x-spend >= arr[startIndex][0] { //原价可以买得起
			dfs(arr, spend+arr[startIndex][0], discount, startIndex+1, cnt+1)
		}

		// 尝试优惠买
		if discount < y && x-spend >= arr[startIndex][1] {
			dfs(arr, spend+arr[startIndex][1], discount+1, startIndex+1, cnt+1)
		}
		// 不买
		dfs(arr, spend, discount, startIndex+1, cnt)

	}

	dfs(goods, 0, 0, 0, 0)
	fmt.Printf("%d %d", maxCnt, minSpend)

}
