package main

func minOperations(n int) int {
	result := 0
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
			n >>= 1
			continue
		}
		n >>= 1
		if count == 0 {
			continue
		}
		// 超过1个的1转成加1，剩余一个1
		// 只有1个1，就只加一消掉
		result++
		if count == 1 {
			count = 0
		} else {
			count = 1
		}
	}
	if count > 2 {
		result += 2
	} else {
		result += count
	}

	return result
}
