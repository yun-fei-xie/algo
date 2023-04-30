package _1_array

func setZeroes(matrix [][]int) {
	maxRow := len(matrix)
	maxCol := len(matrix[0])
	// 用一个标记数组，标记哪一行、哪一列需要置零 最后统一处理。
	// 不能边标记边处理，会影响到后面的标记，导致标记出错。
	tag := make([][2]int, 0)

	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			if matrix[i][j] == 0 {
				tag = append(tag, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(tag); i++ {
		setZeroRowAndCol(matrix, tag[i][0], tag[i][1])
	}

}

func setZeroRowAndCol(matrix [][]int, row, col int) {
	maxRow := len(matrix)
	maxCol := len(matrix[0])
	// 行
	for i := 0; i < maxCol; i++ {
		matrix[row][i] = 0
	}
	for i := 0; i < maxRow; i++ {
		matrix[i][col] = 0
	}

}
