package main

func sumImbalanceNumbers(nums []int) int {
	n := len(nums)
	res := 0
	for i, v := range nums {
		l, r := i-1, i+1
		ll := -1
		for ; l >= 0 && (nums[l] != v && nums[l] != v-1); l-- {
			if ll == -1 && nums[l] < v {
				ll = l
			}
		}
		rl := -1
		for ; r < n && nums[r] != v-1; r++ {
			if rl == -1 && nums[r] < v {
				rl = l
			}
		}
		if ll > 0 {
			res += (ll - l) * (r - i)
		}
		if rl > 0 {
			res += (i - l) * (rl - i)
		}
	}
	return res
}
