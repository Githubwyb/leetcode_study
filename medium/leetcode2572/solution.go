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
	f[0] = 1 // 空背包代表一种方案，所有元素要放到空背包中
	for _, v := range nums {
		target := getTarget(v, p)
		if target == -1 {
			continue
		}
		// 这个元素可以放到背包，找现在所有的方案中可以放进去的
		for i, v := range f {
			if v != 0 && i&target == 0 {
				v1 := i | target
				f[v1] = (f[v1] + v) % mod
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans = (ans + v) % mod
	}
	return ans - 1
}

func squareFreeSubsets1(nums []int) int {
	p := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var mod int = 1e9 + 7
	// 构建背包
	f := make(map[int]int)
	f[0] = 1 // 空背包代表一种方案，所有元素要放到空背包中
	for _, v := range nums {
		target := getTarget(v, p)
		if target == -1 {
			continue
		}
		// 这个元素可以放到背包，找现在所有的方案中可以放进去的
		for i, v := range f {
			if i&target == 0 {
				v1 := i | target
				f[v1] = (f[v1] + v) % mod
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans = (ans + v) % mod
	}
	// 最终返回的不包含空背包方案
	return ans - 1
}
