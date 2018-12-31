package algo

import "sort"

func isContinues(nums []int) bool {
	if len(nums) < 5 {
		return false
	}

	sort.Ints(nums)

	countOf0 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countOf0++
		}
	}
	small := countOf0
	big := small + 1
	diff := 0
	for big < 5 {
		if nums[small] == nums[big] {
			return false
		}
		diff += nums[big] - nums[small] - 1
		small = big
		big++
	}
	if diff > countOf0 {
		return false
	}
	return true
}
