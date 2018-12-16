package algo

func findInTwoDimensionArray(nums [][]int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	columns := len(nums[0]) - 1
	rows := len(nums) - 1
	row := 0
	column := columns
	for row <= rows && column >= 0 {
		if nums[row][column] == target {
			return true
		} else if nums[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}
