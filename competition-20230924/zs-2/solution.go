package main

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	// 枚举每个位置为峰值点，找最大
	var res int64
	for i, v := range maxHeights {
		height := int64(v)

		maxH := v
		for j := i - 1; j >= 0; j-- {
			if maxHeights[j] < maxH {
				maxH = maxHeights[j]
			}
			height += int64(maxH)
		}

		maxH = v
		for j := i + 1; j < n; j++ {
			if maxHeights[j] < maxH {
				maxH = maxHeights[j]
			}
			height += int64(maxH)
		}

		if res < height {
			res = height
		}
	}
	return res
}
