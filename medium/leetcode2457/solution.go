package main

func makeIntegerBeautiful(n int64, target int) int64 {
	curSum := 0
	nums := make([]int, 0, 12)
	for n != 0 {
		t := int(n % 10)
		curSum += t
		nums = append(nums, t)
		n = n / 10
	}

	if curSum <= target {
		return 0
	}

	var result int64 = 0
	aa := 1
	for _, v := range nums {
		// 先假设每一位都加到9，最终返回结果要加一
		result += int64(9-v) * int64(aa)
		curSum -= v
		// 排除一位，但是高位要进一，所以判断小于即可
		if curSum < target {
			return result + 1
		}
		aa *= 10
	}
	return result + 1
}
