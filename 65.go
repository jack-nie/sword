package algo

func add(a, b int) int {
	sum, carry := 0, 0
	for {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
		if b == 0 {
			break
		}
	}
	return sum
}
