package main

func beautifulSplits(nums []int) int {
	n := len(nums)
	isPrefix := func(i, j int) bool {
		i1, i2 := i, j
		for i1 < j {
			if nums[i1] != nums[i2] {
				return false
			}
			i1++
			i2++
		}
		return true
	}
	// 先按照nums1是nums2的前缀查找
	res := 0
	a := map[int]bool{}
	for i := 1; i < (n+1)/2; i++ {
		if isPrefix(0, i) {
			res += n - i*2
			a[i] = true
		}
	}
	// 按照nums2是nums3的前缀查找
	for i := 1; i < n-1; i++ {
		for j := i + 1; (a[i] && j-i < i) || (!a[i] && j < (n-i)/2+i+1); j++ {
			if isPrefix(i, j) {
				if a[i] && j-i >= i {
					continue
				}
				res++
			}
		}
	}

	return res
}
