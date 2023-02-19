package main

// 将x根据p转成对应的二进制bit数
func getTarget(x int, p []int) int {
	result := 0
	for i, t := 0, 1; i < len(p) && x >= p[i]; {
		if x%p[i] != 0 {
			i, t = i+1, t<<1
			continue
		}
		if result&t != 0 {
			return -1
		}
		result |= t
		x /= p[i]
	}
	return result
}

func squareFreeSubsets(nums []int) int {
	p := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var mod int = 1e9 + 7
	n := 1 << len(p)
	f := make([]int, n)
	f[0] = 1
	for _, v := range nums {
		target := getTarget(v, p)
		if target == -1 {
			continue
		}
		// 找整个空间中谁可以乘以v，可以乘的就加上除v之后的数量
		for i := 0; i < n; i++ {
			// i可以由某个数乘以target得到才能参与运算
			if i&target == target {
				f[i] = (f[i] + f[i^target]) % mod
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans = (ans + v) % mod
	}
	return ans - 1
}
