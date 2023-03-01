# leetcode 27 移除元素

## 题目链接

https://leetcode.cn/problems/remove-element/


## 解题思路

如果使用额外的空间，感觉很容易，对原树组件进行一次遍历，将val!=target
的元素放入新数组即可。

但是不实用额外的空间，可以想到的就是后位前移。
用一个count记录当前val!=target的值应该前移多少位。
而count表示当前元素前面出现过多少次val==target
例如：如果当前元素前面出现过2次val==target，那么 nums[i-2] = nums[i]

如果当前val==target 指针直接向后移动一位; removedCount ++
如果当前val!=target nums[i-removeCount] = nums[i] ; resCount++
最后返回nums[0:resCount] 这样一个slice。



## 解题code

```go

func removeElement(nums []int, val int) int {

	removedCount := 0
	resLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			removedCount++
		} else {
			nums[i-removedCount] = nums[i]
			resLength++
		}
	}

	nums = nums[:resLength]
	return resLength
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	target := 3
	res := removeElement(nums, target)
	fmt.Println(res)
}

```