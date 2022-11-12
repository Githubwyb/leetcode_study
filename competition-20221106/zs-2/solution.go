package main

func maximumSubarraySum(nums []int, k int) int64 {
	numMap := make(map[int]int)
	var result int64 = 0
	var add int64 = 0
	l := 0
	for r, v := range nums {
		add += int64(v)
		index := 0
		var ok bool
		if index, ok = numMap[v]; !ok || index < l {
			// 在窗口内没有找到相等的
			numMap[v] = r
			if r-l == k {
				add -= int64(nums[l])
				l++
			}

			if r-l == k-1 && result < add {
				result = add
			}
			continue
		}
		// 窗口内相等，将l移动到相等的元素右边，result减去中间的值
		numMap[v] = r

		for i := l; i <= index; i++ {
			add -= int64(nums[i])
		}
		l = index + 1
	}
	return result
}
