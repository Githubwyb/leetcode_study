package main

func reverse(num int) int {
	result := 0
	for ; num != 0; num /= 10 {
		result *= 10
		result += num % 10
	}
	return result
}

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i > num/2 && i%10 != 0 {
			continue
		}
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}
