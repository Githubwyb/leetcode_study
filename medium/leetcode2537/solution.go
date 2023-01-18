package main

func countGood(nums []int, k int) int64 {
	l := 0
	countMap := make(map[int]int)
	check := 0
	var result int64
	// 右侧遍历端点
	for _, v := range nums {
		// 增加的对个数是之前的此元素的个数
		check += countMap[v]
		countMap[v]++
		for check-countMap[nums[l]]+1 >= k {
			// 如果减去左端点后减少了对数还是大于等于k的，就左指针右移
			countMap[nums[l]]--
			check -= countMap[nums[l]]
			l++
		}
		if check >= k {
			result += int64(l + 1)
		}
	}
	return result
}
