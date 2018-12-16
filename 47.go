package algo

func findGreatestValue(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxValues []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			left := 0
			up := 0
			if i > 0 {
				left = maxValues[j]
			}
			if j > 0 {
				up = maxValues[j-1]
			}
			maxValues = append(maxValues, max(left, up)+nums[i][j])
		}
	}
	return maxValues[len(nums[0])-1]
}
