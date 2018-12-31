package algo

func reverse(strs []rune, start, end int) {
	if len(strs) == 0 {
		return
	}

	for start < end {
		strs[start], strs[end] = strs[end], strs[start]
		start++
		end--
	}
	return
}

func reverseString(s string) string {
	strs := []rune(s)
	reverse(strs, 0, len(strs)-1)
	start := 0
	end := 0

	for end < len(strs) {
		if strs[start] == ' ' {
			start++
			end++
		} else if strs[end] == ' ' {
			reverse(strs[start:end], 0, end-start-1)
			end++
			start = end
		} else if end == len(strs)-1 {
			reverse(strs[start:end+1], 0, end-start)
			end++
		} else {
			end++
		}
	}
	return string(strs)
}
