package algo

func maxCuttingRobe(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	product := make([]int, n+1)
	product[0] = 0
	product[1] = 1
	product[2] = 2
	product[3] = 3
	max := 0
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			tmp := product[j] * product[i-j]
			if tmp > max {
				max = tmp
			}
			product[i] = max
		}
	}
	return product[n]
}
