package mid

import (
	"fmt"
	"sort"
	"testing"
)

/*
输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
排序= [7,0],[7,1],[6,1],[5,0],[5,2],[4,4]

1. [7,0]
2. [7,0],[7,1]
3. [7,0],[6,1],[7,1]
4. [5,0],[7,0],[6,1],[7,1]
5. [5,0],[7,0],[5,2],[6,1],[7,1]
6. [5,0],[7,0],[5,2],[6,1],[4,4],[7,1]

1. 将身高降序排列-> 身高高的排放在前面
2. 相同身高的元素-> 按照k值降序排列，k值大的放在后面

输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

[h,k] k较大，说明h较小。
*/
func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		} else {
			return people[i][1] < people[j][1]
		}
	})

	for i := 0; i < len(people); i++ {
		// end = i-1  start = people[i][1]
		data := people[i]
		start := people[i][1]
		end := i - 1
		shift(people, start, end)
		people[start] = data
	}

	return people

}

func shift(arr [][]int, start int, end int) {
	if start < 0 || end < 0 {
		return
	}
	// arr[p...end] 整体向后移动一位
	for i := end; i >= start; i-- {
		arr[i+1] = arr[i]
	}
}

func TestReconstructQueue(t *testing.T) {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
