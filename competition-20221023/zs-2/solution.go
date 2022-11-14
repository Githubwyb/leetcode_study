package main

func mgcd(a int, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func subarrayGCD(nums []int, k int) int {
	gcdArray := make([]int, len(nums))
	for i := range nums {
		if nums[i]%k == 0 {
			gcdArray[i] = nums[i] / k
		}
	}

	result := 0
	for i := range nums {
		if gcdArray[i] == 0 {
			continue
		}

		tmp := gcdArray[i]
		if tmp == 1 {
			result++
		}
		for j := i + 1; j < len(nums); j++ {
			if gcdArray[j] == 0 {
				break
			}

			tmp = mgcd(tmp, gcdArray[j])
			if tmp == 1 {
				result++
			}
		}
	}
	return result
}
