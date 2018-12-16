package algo

func findGreatestSumOfSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var sum, greatestSum int
	for i := 0; i < len(nums); i++ {

		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > greatestSum {
			greatestSum = sum
		}
	}
	return greatestSum
}
