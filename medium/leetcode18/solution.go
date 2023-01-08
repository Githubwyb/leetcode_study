package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)
	// 然后前两个用循环套
	// 第一个元素首先不能和随后的三个相加已经大于target了，毕竟排了序，大于后面随便选都不可能等于target
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 下一轮循环时第一个元素必须选和当前不一样的，否则就重复了
		// 并且不能和最大的三个相加还比target小，否则同样无法满足选四个相加等于target
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target || (i > 0 && nums[i] == nums[i-1]) {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target || (j > i+1 && nums[j] == nums[j-1]) {
				continue
			}
			// 剩余两个数使用双指针找
			for l, r, t := j+1, n-1, target-nums[i]-nums[j]; r > l; {
				tmp := nums[r] + nums[l] - t
				if tmp == 0 {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
				}
				if tmp >= 0 {
					for r--; r > l && nums[r] == nums[r+1]; r-- {
					}
				}
				if tmp <= 0 {
					for l++; r > l && nums[l] == nums[l-1]; l++ {
					}
				}
			}
		}
	}
	return result
}
