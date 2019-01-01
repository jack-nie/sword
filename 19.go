package algo

var patternIndex, strIndex int

func match(str string, pattern string) bool {
	if len(str) == 0 || len(pattern) == 0 {
		return false
	}
	byteStr := []byte(str)
	bytePattern := []byte(pattern)
	strIndex := 0
	patternIndex := 0
	return matchCore(byteStr, bytePattern, strIndex, patternIndex)
}

func matchCore(byteStr, bytePattern []byte, strIndex, patternIndex int) bool {
	if len(byteStr) == strIndex && len(bytePattern) == patternIndex {
		return true
	}

	if len(byteStr) != strIndex && len(bytePattern) == patternIndex {
		return false
	}
	if patternIndex+1 < len(bytePattern) && bytePattern[patternIndex] == '*' {
		if (strIndex != len(byteStr) && byteStr[strIndex] == bytePattern[patternIndex]) ||
			(bytePattern[patternIndex] == '.' && len(byteStr) != strIndex) {
			return matchCore(byteStr, bytePattern, strIndex, patternIndex+2) ||
				matchCore(byteStr, bytePattern, strIndex+1, patternIndex+2) ||
				matchCore(byteStr, bytePattern, strIndex+1, patternIndex)
		}
		return matchCore(byteStr, bytePattern, strIndex, patternIndex+2)
	}
	if (strIndex != len(byteStr) && byteStr[strIndex] == bytePattern[patternIndex]) ||
		(bytePattern[patternIndex] == '.' && len(byteStr) != strIndex) {
		return matchCore(byteStr, bytePattern, strIndex+1, patternIndex+1)
	}
	return false
}
