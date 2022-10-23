package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	keyIndexs := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxK || nums[i] < minK {
			nums[i] = 4
			keyIndexs = append(keyIndexs, i)
			continue
		}
		if nums[i] < maxK && nums[i] > minK {
			nums[i] = 0
			continue
		}
		check := 0
		if nums[i] == minK {
			check = 1
		}
		if nums[i] == maxK {
			check |= 2
		}
		nums[i] = check
		keyIndexs = append(keyIndexs, i)
	}

	var result int64 = 0
	for i, v := range keyIndexs {
		index := v
		check := nums[index]
		if check == 4 {
			continue
		}
		for j := i; j < len(nums) && nums[j] != 4; j++ {
			check |= nums[j]
			if check == 3 {
				result++
			}
		}
	}
	return result
}
