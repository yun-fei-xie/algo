package unionFind_test

import (
	"fmt"
	"sort"
	"testing"
)

/*
721. 账户合并
https://leetcode.cn/problems/accounts-merge/description/

给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。
现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。
合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。

方法：对每一个节点进行编号，两个节点是否可以进行合并取决于这两个节点背后的邮箱是否有重叠。
*/
func accountsMerge(accounts [][]string) [][]string {
	// 合并
	lenAccount := len(accounts)
	uf := InitUnionFind2(lenAccount)
	for i := 0; i < lenAccount; i++ {
		for j := i + 1; j < lenAccount; j++ {
			if canMerge(accounts[i], accounts[j]) {
				uf.unionElements(i, j)
			}
		}
	}

	// 收集结果 同一个root放在一组
	accountMap := make(map[int][]string)
	for i := 0; i < lenAccount; i++ {
		root := uf.find(i)
		accountMap[root] = append(accountMap[root], accounts[i][1:]...)
	}
	// 对同一组中的邮箱进行去重
	ans := make([][]string, 0)
	for root, account := range accountMap {
		// 首先拿到root的用户名
		person := make([]string, 0)
		person = append(person, accounts[root][0])
		// hash表去重
		set := make(map[string]struct{})
		for i := 0; i < len(account); i++ {
			set[account[i]] = struct{}{}
		}
		// 重hash表中收集答案
		for email, _ := range set {
			person = append(person, email)
		}
		// 排序的时候，第一个string是用户名，不需要进行排序
		sort.Strings(person[1:])
		ans = append(ans, person)

	}
	return ans
}

// 判断两个账户是否属于同一个人 第一位是用户名 不需要参与判断
func canMerge(account1 []string, account2 []string) bool {
	for i := 1; i < len(account1); i++ {
		for j := 1; j < len(account2); j++ {
			if account1[i] == account2[j] {
				return true
			}
		}
	}
	return false
}

func TestAccountsMerge(t *testing.T) {
	fmt.Println(accountsMerge(
		[][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"Mary", "mary@mail.com"}}))
}
