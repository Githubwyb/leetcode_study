package main

func countSubarrays(nums []int, k int) int {
	kIs := []int{}
	for i := range nums {
		if nums[i] == k {
			kIs = append(kIs, i)
		}
	}
	result := 0
	for _, v := range kIs {
		result++
		recordMap := make(map[int]int)
		tmp := 0
		recordMap[0] = 1
		for i := v - 1; i >= 0; i-- {
			if nums[i] > k {
				tmp++
			} else {
				tmp--
			}
			if tmp == 0 || tmp == 1 {
				result++
			}
			key := tmp
			if _, ok := recordMap[key]; !ok {
				recordMap[key] = 0
			}
			recordMap[key]++
		}

		tmp = 0
		for i := v + 1; i < len(nums); i++ {
			var key int
			if nums[i] > k {
				tmp++
			} else {
				tmp--
			}
			key = 1 - tmp
			if v, ok := recordMap[key]; ok {
				result += v
			}
			key = 0 - tmp
			if v, ok := recordMap[key]; ok {
				result += v
			}
		}
	}
	return result
}
