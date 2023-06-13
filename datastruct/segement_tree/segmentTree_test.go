package segement_tree_test

import (
	"algo/datastruct/segement_tree"
	"fmt"
	"sort"
	"testing"
)

func TestSegmentTree_Query(t *testing.T) {

	arr := []int{1564900000, 1564900005, 1564900008, 1564900007, 1564900009}
	sort.Ints(arr)
	segTree := segement_tree.BuildSegmentTree(arr)
	q1 := segTree.Query(0, 9999999999)
	q2 := segTree.Query(1564900008, 1564900008)
	q3 := segTree.Query(1564900006, 1564900006)
	q4 := segTree.Query(1564900007, 1564900009)
	q5 := segTree.Query(1564900000, 1564900007)
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)
	fmt.Println(q5)
}

func TestSegment2(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}

	segmentTree := segement_tree.NewSegmentTree(arr, func(i, j int) int {
		return i + j
	})
	fmt.Println(segmentTree.Query(0, 7))
	fmt.Println(segmentTree.Query(4, 6))

	segmentTree.Update(0, 100)
	fmt.Println(segmentTree.Query(0, 7))

}
