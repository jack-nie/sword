package algo

func findNumsOf1Between1AndN(n int) (count int) {
	count = 0
	for i := 1; i <= n; i++ {
		count += numberOf1(i)
	}
	return
}

func numberOf1(n int) int {
	var num int
	for n != 0 {
		if n%10 == 1 {
			num++
		}
		n /= 10
	}
	return num
}
