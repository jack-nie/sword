package algo

func replaceSpace(str string) string {
	if len(str) == 0 {
		return str
	}

	numbOfSpace := countSpace(str)
	length := len(str) + 2*numbOfSpace
	result := make([]rune, length)
	j := 0
	for _, r := range str {
		if r == ' ' {
			result[j] = '%'
			j++
			result[j] = '2'
			j++
			result[j] = '0'
			j++
		} else {
			result[j] = r
			j++
		}
	}
	return string(result)
}

func countSpace(str string) int {
	count := 0
	for _, r := range str {
		if r == ' ' {
			count++
		}
	}
	return count
}
