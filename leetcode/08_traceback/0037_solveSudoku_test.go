package _8_traceback

// 图的遍历
func solveSudoku(board [][]byte) {

	var traceback func(board [][]byte) bool
	traceback = func(board [][]byte) bool {
		// 递归终止条件：所有的.都被数字替换掉
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == '.' { // 如果这个地方是一个点 尝试放置一个数组
					for k := '1'; k <= '9'; k++ {
						if isVaild(i, j, byte(k), board) { // 每一个位置都合法才会向下进行递归
							board[i][j] = byte(k)
							if traceback(board) == true {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false // 如果当前.从1...9都不行，那么这层递归也不需要继续进行 直接返回false
				}
			}
		}
		return true
	}
	traceback(board)
}

func isVaild(row int, col int, value byte, board [][]byte) bool {
	// 判断行
	for i, l := 0, len(board[0]); i < l; i++ {
		if board[row][i] == value {
			return false
		}
	}
	// 判断列
	for i, l := 0, len(board); i < l; i++ {
		if board[i][col] == value {
			return false
		}
	}
	//判断方格
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == value {
				return false
			}
		}
	}
	return true
}
