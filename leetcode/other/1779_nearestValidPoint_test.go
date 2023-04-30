package other

import (
	"fmt"
	"math"
	"testing"
)

func nearestValidPoint(x int, y int, points [][]int) int {

	var minDist = math.MaxInt32
	var minIndex = -1
	var length = len(points)
	for i := 0; i < length; i++ {
		if points[i][0] == x || points[i][1] == y {
			d := dist(x, y, points[i][0], points[i][1])
			if d < minDist {
				minDist = d
				minIndex = i
			}
		}
	}
	return minIndex
}

func dist(x1, y1, x2, y2 int) int {
	var abs func(i, j int) int
	abs = func(i, j int) int {
		if i > j {
			return i - j
		}
		return j - i
	}

	return abs(x1, x2) + abs(y1, y2)
}

func TestNearestValidPoint(t *testing.T) {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))
}
