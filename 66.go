package algo

func multiply(nums1, nums2 []int) []int {
	length1, length2 := len(nums1), len(nums2)
	if length1 == length2 && length2 > 1 {
		nums2[0] = 1
		for i := 1; i < length1; i++ {
			nums2[i] = nums2[i-1] * nums1[i-1]
		}
		tmp := 1
		for i := length1 - 2; i >= 0; i-- {
			tmp *= nums1[i+1]
			nums2[i] *= tmp
		}
	}
	return nums2

}
