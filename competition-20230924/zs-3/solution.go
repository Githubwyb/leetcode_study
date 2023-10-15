package main

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	// 计算后缀和
	sufMax := make([]int64, n+1)
	st := make([]int, 0, n) // 只存下标

	for i := range maxHeights {
		idx := n - 1 - i
		v := maxHeights[idx]
		if i == 0 {
			st = append(st, idx)
			sufMax[idx] = int64(v)
			continue
		}

		// 找到栈中第一个小于等于当前值的位置
		j := len(st) - 1
		for ; j >= 0 && v < maxHeights[st[j]]; j-- {
		}

		if j < 0 {
			// 都比当前值小，直接全部赋值成当前值
			sufMax[idx] = int64(i+1) * int64(v)
			st = append(st[:0], idx)
		} else {
			// 这个较小值的位置的最大值加上剩下的全部赋值当前值
			sufMax[idx] = sufMax[st[j]] + int64(v)*int64(st[j]-idx)
			st = append(st[:j+1], idx)
		}
	}

	st = st[:0]
	res := sufMax[0]
	preMax := make([]int64, n)
	for i, v := range maxHeights {
		if i == 0 {
			st = append(st, i)
			preMax[i] = int64(v)
		} else {
			// 找到栈中第一个小于等于当前值的位置
			j := len(st) - 1
			for ; j >= 0 && v < maxHeights[st[j]]; j-- {
			}

			if j < 0 {
				// 都比当前值小，直接全部赋值成当前值
				preMax[i] = int64(i+1) * int64(v)
				st = append(st[:0], i)
			} else {
				// 这个较小值的位置的最大值加上剩下的全部赋值当前值
				preMax[i] = preMax[st[j]] + int64(v)*int64(i-st[j])
				st = append(st[:j+1], i)
			}
		}

		x := sufMax[i+1] + preMax[i]
		if x > res {
			res = x
		}
	}

	return res
}
