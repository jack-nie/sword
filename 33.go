package algo

func verifySequenceOfBST(sequence []int, length int) bool {
	if len(sequence) == 0 || length <= 0 {
		return false
	}
	root := sequence[length-1]
	i := 0
	for ; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}
	j := i
	for ; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}
	left := true
	if i > 0 {
		left = verifySequenceOfBST(sequence, i)
	}
	right := true
	if j < length-1 {
		right = verifySequenceOfBST(sequence[i:], length-i-1)
	}
	return left && right

}
