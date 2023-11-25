package main

func minGroupsForValidAssignment(nums []int) int {
	intMap := map[int]int{}
	for _, v := range nums {
		intMap[v]++
	}

	checkValid := func(num, minDiv int) int {
		if minDiv > num {
			return 0
		}
		n, y := num/minDiv, num%minDiv
		if y > n {
			return 0
		}
		i := 1
		for ; i < n && y+i*minDiv <= n-i; i++ {
		}
		return n - i + 1
	}

	nums = nums[:0]
	nums1 := []int{}
	first := true
	for _, v := range intMap {
		if first {
			first = false
			for i := v; i > 0; i-- {
				if n := checkValid(v, i); n != 0 {
					nums = append(nums, i)
				}
			}
			continue
		}
		nums1 = nums1[:0]
		for _, v1 := range nums {
			if n := checkValid(v, v1); n != 0 {
				nums1 = append(nums1, v1)
			}
		}
		nums = nums[:len(nums1)]
		copy(nums, nums1)
	}

	div := nums[0]
	res := 0
	for _, v := range intMap {
		res += checkValid(v, div)
	}
	return res

}
