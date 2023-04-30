package jianzhi_offer

/*
numbers = [3,4,5,1,2]
从左到右遍历，符合numbers[i] < numbers[i-1]的第一个数字就是最小的数字
如果数组全部都是升序的，那么数组的第一个元素就是最小的元素。

*/

func minArray(numbers []int) int {
	length := len(numbers)
	for i := 1; i < length; i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}
