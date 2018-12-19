package algo

func isPopOrder(push, pop []int, length int) bool {
	var flag bool
	var pushIndex, popIndex int
	var stack []int
	if len(push) != 0 && len(pop) != 0 && length > 0 {
		for popIndex < length {
			for len(stack) == 0 || stack[len(stack)-1] != pop[popIndex] {
				if pushIndex == length {
					break
				}
				stack = append(stack, push[pushIndex])
				pushIndex++
			}
			if stack[len(stack)-1] != pop[popIndex] {
				break
			}
			stack = stack[:len(stack)-1]
			popIndex++
		}
		if len(stack) == 0 && popIndex == length {
			flag = true
		}
	}
	return flag
}
