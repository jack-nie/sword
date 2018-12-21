package algo

func fibnacci(n int) int {
	if n < 0 || n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fibNMinusOne := 0
	fibNMinusTwo := 1
	var fibN int
	for i := 2; i <= n; i++ {
		fibN = fibNMinusOne + fibNMinusTwo
		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
	}
	return fibN
}
