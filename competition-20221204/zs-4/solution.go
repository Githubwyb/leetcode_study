package main

func countSubarrays(nums []int, k int) int {
	result := 0
	v := len(nums)
	tmp := 0
	recordMap := make(map[int]int)
	recordMap[0] = 1 // 第一位是0
	for i := range nums {
		if nums[i] > k {
			tmp++
		} else if nums[i] < k {
			tmp--
		} else {
			v = i
		}
		// 每个i计算的都是下一位的值，所以不需要到中位数所在的位置
		if i < v {
			if _, ok := recordMap[tmp]; !ok {
				recordMap[tmp] = 0
			}
			recordMap[tmp]++
		} else {
			if count, ok := recordMap[tmp]; ok {
				result += count
			}
			if count, ok := recordMap[tmp-1]; ok {
				result += count
			}
		}
	}
	return result
}

func countSubarrays1(nums []int, k int) int {
	// 1. 先找中位数的位置
	v := 0
	for i := range nums {
		if nums[i] == k {
			v = i
			break
		}
	}

	// 2. 向左开始进行累加，使用map进行存储
	recordMap := make(map[int]int)
	recordMap[0] = 1 // 中位数的位置本身就是0
	tmp := 0
	for i := v - 1; i >= 0; i-- {
		if nums[i] > k {
			tmp--
		} else {
			tmp++
		}
		if _, ok := recordMap[tmp]; !ok {
			recordMap[tmp] = 0
		}
		recordMap[tmp]++
	}

	// 向右开始查找
	result := 0
	tmp = 0
	for i := v; i < len(nums); i++ {
		if nums[i] > k {
			tmp++
		} else if nums[i] < k {
			tmp--
		}
		// 左侧小于-左侧大于 = 右侧大于-右侧小于
		if count, ok := recordMap[tmp]; ok {
			result += count
		}
		// 左侧小于-左侧大于+1 = 右侧大于-右侧小于
		if count, ok := recordMap[tmp-1]; ok {
			result += count
		}
	}
	return result
}
