package main

func continuousSubarrays(nums []int) int64 {
	// 枚举左端点，找到不间断的子数组
	n := len(nums)
	var res int64
	intMap := make(map[int]int)
	r := 0
	for l, x := range nums {
		min, max := x, x
		for k := range intMap {
			if k < min {
				min = k
			}
			if k > max {
				max = k
			}
		}

		// 找窗口右端点，从上一个右端点向右遍历
		for ; r < n; r++ {
			v := nums[r]
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
			if max-min > 2 {
				break
			}
			intMap[v]++
		}
		res += int64(r - l)
		if intMap[x]--; intMap[x] == 0 {
			delete(intMap, x)
		}
	}
	return res
}
