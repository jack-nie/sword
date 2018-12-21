package algo

func numsOf1(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

func numsOf1Two(n int) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
