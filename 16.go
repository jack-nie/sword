package algo

func power(base float64, exp int) float64 {
	if doubleEqual(base, 0.0) {
		return 0.0
	}
	absExp := uint(exp)
	if exp < 0 {
		absExp = uint(-exp)
	}
	result := powerWithUnsignedExp(base, absExp)
	if exp < 0 {
		result = 1 / result
	}
	return result
}

func powerWithUnsignedExp(base float64, exp uint) float64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	result := powerWithUnsignedExp(base, exp>>1)
	result *= result
	if exp&0x1 == 1 {
		result *= base
	}
	return result
}
