package algo

func findGreatestSumOfSubArray(nums []int) (result int) {
	result = 0
	if len(nums) == 0 {
		return
	}
	var sum int
	for i := 0; i < len(nums); i++ {

		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > result {
			result = sum
		}
	}
	return
}
