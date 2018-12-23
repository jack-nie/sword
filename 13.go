package algo

func movingCount(threshold, rows, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	return movingCountCore(threshold, rows, cols, 0, 0, &visited)
}

func movingCountCore(threshold, rows, cols, row, col int, visited *[][]bool) int {
	count := 0
	if check(threshold, rows, cols, row, col, visited) {
		(*visited)[row][col] = true
		count = 1 + movingCountCore(threshold, rows, cols, row+1, col, visited) +
			movingCountCore(threshold, rows, cols, row-1, col, visited) +
			movingCountCore(threshold, rows, cols, row, col+1, visited) +
			movingCountCore(threshold, rows, cols, row, col-1, visited)
	}
	return count
}

func check(threshold, rows, cols, row, col int, visited *[][]bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		getDigitSum(row)+getDigitSum(col) < threshold && !(*visited)[row][col] {
		return true
	}
	return false
}

func getDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
