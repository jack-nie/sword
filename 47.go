package algo

func findGreatestValue(nums [][]int) (result int) {
	result = 0
	if len(nums) == 0 {
		return
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
	result = maxValues[len(nums[0])-1]
	return
}
