package main

func sumOfTheDigitsOfHarshadNumber(n int) int {
	res := 0
	tmp := n
	for n != 0 {
		res += n % 10
		n /= 10
	}

	if tmp%res == 0 {
		return res
	}
	return -1
}
