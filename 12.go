package algo

func routeInMatrix(matrix [][]string, strs []string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}

	pathLength := 0
	rows := len(matrix)
	cols := len(matrix[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasPath(matrix, rows, cols, row, col, strs, &pathLength, &visited) {
				return true
			}
		}
	}
	return false
}

func hasPath(matrix [][]string, rows, cols, row, col int, strs []string, pathLength *int, visited *[][]bool) bool {
	if len(strs) == *pathLength {
		return true
	}

	flag := false
	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row][col] == strs[*pathLength] && !(*visited)[row][col] {
		(*pathLength)++
		(*visited)[row][col] = true
		flag = hasPath(matrix, rows, cols, row+1, col, strs, pathLength, visited) ||
			hasPath(matrix, rows, cols, row-1, col, strs, pathLength, visited) ||
			hasPath(matrix, rows, cols, row, col+1, strs, pathLength, visited) ||
			hasPath(matrix, rows, cols, row, col-1, strs, pathLength, visited)
		if !flag {
			(*pathLength)--
			(*visited)[row][col] = false
		}

	}
	return flag
}
