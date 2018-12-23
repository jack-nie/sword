package algo

import "fmt"

func print1ToMaxNOfDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]byte, n)
	for i := 0; i < 10; i++ {
		number[0] = byte(i) + '0'
		print1ToMaxNOfDigitsRecursively(number, n, 0)
	}
	return
}

func print1ToMaxNOfDigitsRecursively(number []byte, length, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}
	for i := 0; i < 10; i++ {
		number[index+1] = byte(i) + '0'
		print1ToMaxNOfDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []byte) {
	length := len(number)
	isBeginning0 := true
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Printf("%s", string(number[i]))
		}
	}
	fmt.Println()
}
