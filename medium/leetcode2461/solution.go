package main

func maximumSubarraySum(nums []int, k int) int64 {
	numMap := make(map[int]int) // value为最近一次出现的索引
	var res int64 = 0
	var add int64 = 0
	l := 0 // 左指针，包含
	for r, v := range nums {
		add += int64(v)
		index, ok := numMap[v]
		if ok && index >= l {
			// 找到相等的，并且就在范围内，将左侧左移到index右边
			for i := l; i <= index; i++ {
				add -= int64(nums[i])
			}
			l = index + 1
		} else if r-l == k-1 {
			if add > res {
				res = add
			}
			add -= int64(nums[l])
			l++
		}
		numMap[v] = r
	}
	return res
}
