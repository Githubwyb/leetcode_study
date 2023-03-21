package main

func getCount(result map[int]int, i, k int, nums []int, cb func()) {
	if i == len(nums) {
		if len(result) > 0 {
			cb()
		}
		return
	}
	getCount(result, i+1, k, nums, cb)
	v := nums[i]
	if result[v-k] > 0 || result[v+k] > 0 {
		return
	}
	result[v]++
	getCount(result, i+1, k, nums, cb)
	result[v]--
	if result[v] == 0 {
		delete(result, v)
	}
}

func beautifulSubsets(nums []int, k int) int {
	result := make(map[int]int)
	res := 0
	getCount(result, 0, k, nums, func() {
		res++
	})
	return res
}
