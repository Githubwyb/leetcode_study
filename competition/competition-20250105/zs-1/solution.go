package main

func gcd(a int, b int) int {
	// 当a为最大公约数时，计算后a = 0，b = a
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func maxLength(nums []int) int {
	ans := 2
	for i, v := range nums {
		l := 1
		for _, v1 := range nums[i+1:] {
			if gcd(v1, v) == 1 {
				l++
				v *= v1
				continue
			}
			break
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
