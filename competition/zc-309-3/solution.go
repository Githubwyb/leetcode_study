package main

func longestNiceSubarray(nums []int) int {
	// 记录每一个数字向后的优雅数量
	niceMap := make([]int, len(nums))
	for i := range nums {
		niceMap[i] = 1
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[i]&nums[j] != 0 {
				break
			}
		}
		if j-i > 1 {
			niceMap[i] = j - i
		}
	}
	// 遍历niceMap查看最大相互优雅数量
	result := 1
	// 倒序遍历，在范围内就++
	for i := len(nums) - 1; i >= result; i-- {
		j := i - 1
		for ; j >= 0; j-- {
			if niceMap[j] < i-j+1 {
				break
			}
		}
		if i-j > result {
			result = i - j
		}
	}
	return result
}
