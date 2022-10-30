package main

func goodIndices(nums []int, k int) []int {
	leftK := k - 1
	rn := 1 // 代表index右侧多少个确认了非递减
	result := []int{}
	for index := range nums {
		if index == 0 {
			continue
		}
		// 下一个循环，上次遍历过的非递增数量-1
		if rn > 1 {
			rn--
		}
		if leftK > 0 {
			// 先找k个非递增元素
			if nums[index] <= nums[index-1] {
				leftK--
			} else {
				// 否则，左指针移到当前
				leftK = k - 1
			}
			continue
		}
		// 剩余的不到k
		if len(nums)-index-1 < k {
			break
		}
		// 查看右侧k个元素是否非递减
		i := rn
		for ; i < k; i++ {
			if nums[index+i+1] < nums[index+i] {
				break
			}
			rn++
		}
		if i == k {
			result = append(result, index)
		}
		if nums[index] > nums[index-1] {
			leftK = k - 1
		}
	}
	return result
}

func goodIndices1(nums []int, k int) []int {
	n := len(nums)
	left := make([]int, n)  // left[i]表示包括i在内，左侧多少满足非递增
	right := make([]int, n) // right[i]表示包括i在内，右侧多少满足非递减
	for i := range nums {
		left[i] = 1
		right[i] = 1
	}

	// 计算left和right
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] <= nums[i-1] {
			left[i] = left[i-1] + 1
		}
		if nums[n-i-1] <= nums[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	// 如果某个元素左侧非递增大于等于k且右侧非递减大于等于k，则满足要求
	result := []int{}
	for i := k; i < n-k; i++ {
		if left[i-1] >= k && right[i+1] >= k {
			result = append(result, i)
		}
	}
	return result
}

func goodIndices2(nums []int, k int) []int {
	n := len(nums)
	right := make([]int, n) // right[i]表示包括i在内，右侧多少满足非递减

	// 计算left
	for i := range nums {
		if i != 0 && nums[n-i-1] <= nums[n-i] {
			right[n-i-1] = right[n-i] + 1
			continue
		}
		right[n-i-1] = 1
	}

	// 如果某个元素左侧非递增大于等于k且右侧非递减大于等于k，则满足要求
	result := []int{}
	left := 1
	for i := 1; i < n-k; i++ {
		if left >= k && right[i+1] >= k {
			result = append(result, i)
		}
		if nums[i] <= nums[i-1] {
			left++
		} else {
			left = 1
		}
	}
	return result
}
