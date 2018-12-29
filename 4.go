package algo

func findInTwoDimensionArray(nums [][]int, target int) (result bool) {
	result = false
	if len(nums) == 0 {
		return
	}

	columns := len(nums[0]) - 1
	rows := len(nums) - 1
	row := 0
	column := columns
	for row <= rows && column >= 0 {
		if nums[row][column] == target {
			result = true
			return
		} else if nums[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return
}
