package counting_sort

import (
	"fmt"
	"testing"
)

/*
计数排序
https://zh.wikipedia.org/wiki/%E8%AE%A1%E6%95%B0%E6%8E%92%E5%BA%8F
https://www.runoob.com/w3cnote/counting-sort.html

计数排序只能处理待排数据是是非负数的情况，因为它用数组下标表示元素的值。
为了能够处理这种情况，可以将下标做偏移处理。
比如处理[-5,5] 这个范围的计数排序，那么可以将数组空间开为[0,10]
也就是说，将-5，-4，-3，-2，-1 映射到0，1，2，3，4 （下标都减掉最小值）

[-3,5] -> [0,8]

[6 , 10] -> [0 , 4]  本质还是根据区间的范围，进行平移


[5,1,1,2,0,0] -> [2,2,1,0,0,1]
*/

func countingSort2(arr []int, minValue int, maxValue int) []int {
	counting := make([]int, maxValue-minValue+1)

	for i := 0; i < len(arr); i++ {
		counting[arr[i]-minValue]++ // 这里减去了最小值
	}

	sortedIndex := 0
	for j := 0; j < len(counting); j++ {
		for counting[j] > 0 {
			arr[sortedIndex] = j + minValue // 这里把最小值加回来
			sortedIndex++
			counting[j]--
		}
	}
	return arr
}

func countingSort(arr []int, maxValue int) []int {
	var counting = make([]int, maxValue+1) // 假设最大值是maxValue 并且数组中都是非负数

	for i := 0; i < len(arr); i++ {
		counting[arr[i]]++
	}

	sortIndex := 0
	// [5,1,1,2,0,0] -> [2 ,2 , 1, 0 , 0 ,1]  -> 2个0，2个1，1个2 ，0个3...
	// 下标代表要排序的值，value表示这个值在待排序数组中的个数。因此就需要有 counting[j]>0 这个判断。只处理待排序数据。
	for j := 0; j < len(counting); j++ {
		for counting[j] > 0 {
			arr[sortIndex] = j
			sortIndex++
			counting[j]--
		}
	}
	return arr
}

func Test(t *testing.T) {
	fmt.Println(countingSort([]int{5, 1, 1, 2, 0, 0}, 5))
	fmt.Println(countingSort2([]int{5, -1, -1, -2, 0, 0}, -2, 5))

}
