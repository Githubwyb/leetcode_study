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
