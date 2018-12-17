package algo

func printMatrixClockwisely(nums [][]int) []int {
	if len(nums) == 0 {
		return nil
	}
	start := 0
	rows := len(nums)
	cols := len(nums[0])
	result := make([]int, 0)
	for rows > start*2 && cols > start*2 {
		result = append(result, printCore(nums, start, rows, cols)...)
		start++
	}
	return result
}

func printCore(nums [][]int, start, rows, cols int) []int {
	endX := cols - 1 - start
	endY := rows - 1 - start
	result := make([]int, 0)
	for i := start; i < endX+1; i++ {
		result = append(result, nums[start][i])
	}

	if start < endY {
		for i := start + 1; i <= endY; i++ {
			result = append(result, nums[i][endX])
		}
	}
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			result = append(result, nums[endY][i])
		}
	}
	if start < endX && start < endY-1 {
		for i := endY - 1; i >= start+1; i-- {
			result = append(result, nums[i][start])
		}
	}
	return result
}
