package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	chCount := make([]int, 3)
	// 第一个for循环，查看左侧为0，右侧需要有多长
	for r := len(s) - 1; r >= 0; r-- {
		chCount[int(s[r]-'a')]++
	}
	for _, v := range chCount {
		if v < k {
			return -1
		}
	}

	result := len(s)
	r := 0
	// l所在位置还没有被选入，下一个循环被选入，不用考虑最后一个，因为最后一个和r到第一个一样
	for l := 0; l < result; l++ {
		// 右侧对应的字符大于k，说明右侧可以移动
		for r < len(s) && chCount[int(s[r]-'a')] > k {
			chCount[int(s[r]-'a')]--
			r++
		}
		result = min(result, l+len(s)-r)
		chCount[int(s[l]-'a')]++
	}

	return result
}
